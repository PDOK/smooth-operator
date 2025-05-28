package validation

import (
	"testing"

	"github.com/pdok/smooth-operator/model"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

func TestURLsContainsBaseURL(t *testing.T) {
	URL1, _ := model.ParseURL("http://test.com/test")
	URL2, _ := model.ParseURL("http://test.com/other")
	URL3, _ := model.ParseURL("http://test.com/other/path")
	type args struct {
		urls    model.IngressRouteURLs
		baseURL model.URL
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Contains BaseURL",
			args: args{
				urls:    model.IngressRouteURLs{{URL: model.URL{URL: URL1}}, {URL: model.URL{URL: URL2}}},
				baseURL: model.URL{URL: URL1},
			},
			wantErr: false,
		},
		{
			name: "Does not contain BaseURL",
			args: args{
				urls:    model.IngressRouteURLs{{URL: model.URL{URL: URL1}}, {URL: model.URL{URL: URL2}}},
				baseURL: model.URL{URL: URL3},
			},
			wantErr: true,
		},
		{
			name: "No ingressRouteURLs",
			args: args{
				urls:    model.IngressRouteURLs{},
				baseURL: model.URL{URL: URL3},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateIngressRouteURLsContainsBaseURL(tt.args.urls, tt.args.baseURL, nil)
			if tt.wantErr {
				assert.NotEmpty(t, err)
			} else {
				assert.Empty(t, err)
			}
		})
	}
}

func TestURLsNotRemoved(t *testing.T) {
	URL1, _ := model.ParseURL("http://test.com/test")
	URL2, _ := model.ParseURL("http://test.com/other")
	URL3, _ := model.ParseURL("http://test.com/other/path")
	type args struct {
		old model.IngressRouteURLs
		new model.IngressRouteURLs
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Same set",
			args: args{
				old: model.IngressRouteURLs{{URL: model.URL{URL: URL1}}, {URL: model.URL{URL: URL2}}},
				new: model.IngressRouteURLs{{URL: model.URL{URL: URL1}}, {URL: model.URL{URL: URL2}}},
			},
			wantErr: false,
		},
		{
			name: "Same set (both empty)",
			args: args{
				old: model.IngressRouteURLs{},
				new: model.IngressRouteURLs{},
			},
			wantErr: false,
		},
		{
			name: "One URL removed",
			args: args{
				old: model.IngressRouteURLs{{URL: model.URL{URL: URL1}}, {URL: model.URL{URL: URL2}}},
				new: model.IngressRouteURLs{{URL: model.URL{URL: URL1}}},
			},
			wantErr: true,
		},
		{
			name: "One URL added",
			args: args{
				old: model.IngressRouteURLs{{URL: model.URL{URL: URL1}}, {URL: model.URL{URL: URL2}}},
				new: model.IngressRouteURLs{{URL: model.URL{URL: URL1}}, {URL: model.URL{URL: URL2}}, {URL: model.URL{URL: URL3}}},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			allErrs := field.ErrorList{}
			ValidateIngressRouteURLsNotRemoved(tt.args.old, tt.args.new, &allErrs, nil)
			if tt.wantErr {
				assert.NotEmpty(t, allErrs)
			} else {
				assert.Empty(t, allErrs)
			}
		})
	}
}
