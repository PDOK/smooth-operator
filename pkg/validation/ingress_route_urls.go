package validation

import (
	"fmt"
	"github.com/pdok/smooth-operator/model"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

func ValidateIngressRouteURLsContainsBaseURL(urls model.IngressRouteURLs, baseURL model.URL, path *field.Path) *field.Error {
	if path == nil {
		path = field.NewPath("spec").Child("ingressRouteUrls")
	}

	if len(urls) == 0 {
		return nil
	}

	for _, url := range urls {
		if url.URL.String() == baseURL.String() {
			return nil
		}
	}

	return field.Invalid(path, urls, fmt.Sprintf("must contain baseURL: %s", baseURL))
}

func ValidateIngressRouteURLsNotRemoved(oldURLs, newURLs model.IngressRouteURLs, allErrs *field.ErrorList, path *field.Path) {
	if path == nil {
		path = field.NewPath("spec").Child("ingressRouteUrls")
	}

	for _, url := range oldURLs {
		s := url.URL.String()
		found := false

		for _, newURL := range newURLs {
			if newURL.URL.String() == s {
				found = true
				break
			}
		}

		if !found {
			*allErrs = append(*allErrs, field.Invalid(path, newURLs, fmt.Sprintf("urls cannot be removed: %s", url)))
		}
	}
}
