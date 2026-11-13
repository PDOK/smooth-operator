package reloader

import (
	"flag"
	"io"
	"reflect"
	"testing"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestBindFlag(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		annotation string
	}{
		{name: "default disabled"},
		{name: "bare flag enables auto", args: []string{"--" + FlagName}, annotation: AutoAnnotation},
		{name: "true enables auto", args: []string{"--" + FlagName + "=true"}, annotation: AutoAnnotation},
		{name: "legacy boolean spelling", args: []string{"--" + FlagName + "=1"}, annotation: AutoAnnotation},
		{name: "false disables", args: []string{"--" + FlagName + "=false"}},
		{name: "explicit auto", args: []string{"--" + FlagName + "=auto"}, annotation: AutoAnnotation},
		{name: "explicit search", args: []string{"--" + FlagName + "=search"}, annotation: SearchAnnotation},
		{name: "last value wins", args: []string{"--" + FlagName + "=search", "--" + FlagName + "=false"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() { SetReloaderAnnotations(false) })
			fs := flag.NewFlagSet("test", flag.ContinueOnError)
			BindFlag(fs)
			if err := fs.Parse(tt.args); err != nil {
				t.Fatal(err)
			}
			if IsEnabled() != (tt.annotation != "") {
				t.Fatalf("unexpected enabled state: %t", IsEnabled())
			}
			deployment := &appsv1.Deployment{}
			SetAnnotation(deployment)
			if tt.annotation == "" {
				if len(deployment.Annotations) != 0 {
					t.Fatalf("expected no annotations, got %v", deployment.Annotations)
				}
			} else if !reflect.DeepEqual(deployment.Annotations, map[string]string{tt.annotation: "true"}) {
				t.Fatalf("expected only %s=true, got %v", tt.annotation, deployment.Annotations)
			}
		})
	}
}

func TestBindFlagRejectsInvalidMode(t *testing.T) {
	for _, value := range []string{"", "invalid", "both", "auto,search"} {
		t.Run(value, func(t *testing.T) {
			t.Cleanup(func() { SetReloaderAnnotations(false) })
			fs := flag.NewFlagSet("test", flag.ContinueOnError)
			fs.SetOutput(io.Discard)
			BindFlag(fs)
			if err := fs.Set(FlagName, "search"); err != nil {
				t.Fatal(err)
			}
			if err := fs.Parse([]string{"--" + FlagName + "=" + value}); err == nil {
				t.Fatal("expected invalid mode to be rejected")
			}
			deployment := &appsv1.Deployment{}
			SetAnnotation(deployment)
			if deployment.Annotations[SearchAnnotation] != SearchAnnotationValue {
				t.Fatal("invalid mode changed the previous selection")
			}
		})
	}
}

func TestSetAnnotation(t *testing.T) {
	tests := []struct {
		name        string
		mode        string
		annotations map[string]string
		expected    map[string]string
	}{
		{
			name:     "enabled sets annotation on nil map",
			mode:     "auto",
			expected: map[string]string{AutoAnnotation: AutoAnnotationValue},
		},
		{
			name:        "enabled keeps other annotations",
			mode:        "auto",
			annotations: map[string]string{"foo": "bar"},
			expected:    map[string]string{"foo": "bar", AutoAnnotation: AutoAnnotationValue},
		},
		{
			name:        "disabled removes existing annotation",
			mode:        "false",
			annotations: map[string]string{"foo": "bar", AutoAnnotation: AutoAnnotationValue, SearchAnnotation: SearchAnnotationValue},
			expected:    map[string]string{"foo": "bar"},
		},
		{
			name:     "disabled leaves nil annotations untouched",
			mode:     "false",
			expected: nil,
		},
		{
			name:     "search sets annotation on nil map",
			mode:     "search",
			expected: map[string]string{SearchAnnotation: SearchAnnotationValue},
		},
		{
			name:        "search replaces auto and keeps other annotations",
			mode:        "search",
			annotations: map[string]string{"foo": "bar", AutoAnnotation: AutoAnnotationValue},
			expected:    map[string]string{"foo": "bar", SearchAnnotation: SearchAnnotationValue},
		},
		{
			name:        "auto replaces search and corrects existing value",
			mode:        "auto",
			annotations: map[string]string{"foo": "bar", SearchAnnotation: SearchAnnotationValue, AutoAnnotation: "false"},
			expected:    map[string]string{"foo": "bar", AutoAnnotation: AutoAnnotationValue},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() { SetReloaderAnnotations(false) })
			fs := flag.NewFlagSet("test", flag.ContinueOnError)
			BindFlag(fs)
			if err := fs.Set(FlagName, tt.mode); err != nil {
				t.Fatal(err)
			}
			deployment := &appsv1.Deployment{ObjectMeta: metav1.ObjectMeta{Annotations: tt.annotations}}
			SetAnnotation(deployment)
			SetAnnotation(deployment)
			if got := deployment.GetAnnotations(); !reflect.DeepEqual(got, tt.expected) {
				t.Fatalf("expected %v, got %v", tt.expected, got)
			}
		})
	}
}
