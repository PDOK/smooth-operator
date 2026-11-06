package reloader

import (
	"flag"
	"testing"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestBindFlag(t *testing.T) {
	t.Cleanup(func() { SetReloaderAnnotations(false) })
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	BindFlag(fs)
	if IsEnabled() {
		t.Fatal("expected default to be disabled")
	}
	if err := fs.Parse([]string{"--" + FlagName + "=true"}); err != nil {
		t.Fatal(err)
	}
	if !IsEnabled() {
		t.Fatal("expected flag to enable reloader annotations")
	}
}

func TestSetAnnotation(t *testing.T) {
	tests := []struct {
		name        string
		enabled     bool
		annotations map[string]string
		expected    map[string]string
	}{
		{
			name:     "enabled sets annotation on nil map",
			enabled:  true,
			expected: map[string]string{AutoAnnotation: AutoAnnotationValue},
		},
		{
			name:        "enabled keeps other annotations",
			enabled:     true,
			annotations: map[string]string{"foo": "bar"},
			expected:    map[string]string{"foo": "bar", AutoAnnotation: AutoAnnotationValue},
		},
		{
			name:        "disabled removes existing annotation",
			enabled:     false,
			annotations: map[string]string{"foo": "bar", AutoAnnotation: AutoAnnotationValue},
			expected:    map[string]string{"foo": "bar"},
		},
		{
			name:     "disabled leaves nil annotations untouched",
			enabled:  false,
			expected: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() { SetReloaderAnnotations(false) })
			SetReloaderAnnotations(tt.enabled)
			deployment := &appsv1.Deployment{ObjectMeta: metav1.ObjectMeta{Annotations: tt.annotations}}
			SetAnnotation(deployment)
			got := deployment.GetAnnotations()
			if len(got) != len(tt.expected) {
				t.Fatalf("expected %v, got %v", tt.expected, got)
			}
			for k, v := range tt.expected {
				if got[k] != v {
					t.Fatalf("expected %v, got %v", tt.expected, got)
				}
			}
		})
	}
}
