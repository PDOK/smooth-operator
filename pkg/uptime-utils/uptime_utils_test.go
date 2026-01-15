package uptimeutils

import (
	"reflect"
	"testing"
)

func TestPassUptimeAnnotationsFormCRToIngressRoute(t *testing.T) {
	tests := []struct {
		name                      string
		customResourceAnnotations map[string]string
		ingressRouteAnnotations   map[string]string
		expectedAnnotations       map[string]string
	}{
		{
			name: "should copy uptime annotations",
			customResourceAnnotations: map[string]string{
				"uptime.pdok.nl/id":   "123",
				"uptime.pdok.nl/name": "test-service",
				"other.pdok.nl/foo":   "bar",
			},
			ingressRouteAnnotations: map[string]string{},
			expectedAnnotations: map[string]string{
				"uptime.pdok.nl/id":   "123",
				"uptime.pdok.nl/name": "test-service",
			},
		},
		{
			name: "should append to existing annotations",
			customResourceAnnotations: map[string]string{
				"uptime.pdok.nl/tags": "tag1,tag2",
			},
			ingressRouteAnnotations: map[string]string{
				"existing": "value",
			},
			expectedAnnotations: map[string]string{
				"existing":            "value",
				"uptime.pdok.nl/tags": "tag1,tag2",
			},
		},
		{
			name: "should overwrite existing uptime annotations",
			customResourceAnnotations: map[string]string{
				"uptime.pdok.nl/url": "http://new.url",
			},
			ingressRouteAnnotations: map[string]string{
				"uptime.pdok.nl/url": "http://old.url",
			},
			expectedAnnotations: map[string]string{
				"uptime.pdok.nl/url": "http://new.url",
			},
		},
		{
			name: "should do nothing if no uptime annotations",
			customResourceAnnotations: map[string]string{
				"other.pdok.nl/foo": "bar",
			},
			ingressRouteAnnotations: map[string]string{
				"existing": "value",
			},
			expectedAnnotations: map[string]string{
				"existing": "value",
			},
		},
		{
			name:                      "should handle nil customResourceAnnotations",
			customResourceAnnotations: nil,
			ingressRouteAnnotations:   map[string]string{"foo": "bar"},
			expectedAnnotations:       map[string]string{"foo": "bar"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PassUptimeAnnotationsFromCRToIngressRoute(tt.customResourceAnnotations, tt.ingressRouteAnnotations)
			if !reflect.DeepEqual(tt.ingressRouteAnnotations, tt.expectedAnnotations) {
				t.Errorf("PassUptimeAnnotationsFromCRToIngressRoute() = %v, want %v", tt.ingressRouteAnnotations, tt.expectedAnnotations)
			}
		})
	}
}

func TestSetUptimeAnnotations(t *testing.T) {
	customResourceAnnotations := map[string]string{
		"uptime.pdok.nl/extra": "extra-value",
	}
	id := "1"
	name := "test"
	url := "http://test.com"
	tags := map[string]string{
		"a": "tag1",
		"b": "tag2",
	}

	expected := map[string]string{
		"uptime.pdok.nl/id":    "356a192b7913b04c54574d18c28d46e6395428ab",
		"uptime.pdok.nl/name":  "test",
		"uptime.pdok.nl/url":   "http://test.com",
		"uptime.pdok.nl/tags":  "public-stats,tag1,tag2",
		"uptime.pdok.nl/extra": "extra-value",
	}

	result := GetUptimeAnnotations(customResourceAnnotations, id, name, url, tags)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetUptimeAnnotations() = %v, want %v", result, expected)
	}
}
