package validation

import (
	"testing"

	"github.com/pdok/smooth-operator/model"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

func TestCheckUrlImmutability(t *testing.T) {
	URL1, _ := model.ParseURL("http://test.com/test")
	URL2, _ := model.ParseURL("http://test.com/test")
	URL3, _ := model.ParseURL("http://test.com/other")
	type args struct {
		oldURL model.URL
		newURL model.URL
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "the same URLs",
			args: args{
				oldURL: model.URL{URL: URL1},
				newURL: model.URL{URL: URL2},
			},
			wantErr: false,
		},
		{
			name: "different URLs",
			args: args{
				oldURL: model.URL{URL: URL1},
				newURL: model.URL{URL: URL3},
			},
			wantErr: true,
		},
		{
			name: "nil URLs",
			args: args{
				oldURL: model.URL{},
				newURL: model.URL{},
			},
			wantErr: false,
		},
		{
			name: "old nil, new URL",
			args: args{
				oldURL: model.URL{},
				newURL: model.URL{URL: URL3},
			},
			wantErr: true,
		},
		{
			name: "old URL, new nil",
			args: args{
				oldURL: model.URL{URL: URL1},
				newURL: model.URL{URL: URL3},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			allErrs := field.ErrorList{}
			CheckURLImmutability(tt.args.oldURL, tt.args.newURL, &allErrs, nil)
			if tt.wantErr {
				assert.NotEmpty(t, allErrs)
			} else {
				assert.Empty(t, allErrs)
			}
		})
	}
}
