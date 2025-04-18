package validation

import (
	"context"
	"fmt"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

var validators = map[string]*Validator{}

// ValidateSchema validates incoming data against a OpenAPI specification
func ValidateSchema(yamlString string, crd string) error {
	if val, ok := validators[crd]; ok {
		return val.ValidateCustomResourceYAML(yamlString)
	} else {
		return fmt.Errorf("No schemas found for CRD '%s'", crd)
	}
}

// AddSchema manually add an OpenAPI schema for a CRD
func AddValidator(crd string, schema apiextensions.CustomResourceDefinition) {
	validator, _ := NewValidatorFromCRDs(schema)
	validators[crd] = validator
}

// LoadSchemaForCRD extracts OpenAPI schemas for a specific CRD from a Kubernetes cluster
func LoadSchemasForCRD(cfg *rest.Config, namespace, name string) error {
	crdClientSet, err := clientset.NewForConfig(cfg)
	if err != nil {
		return err
	}

	crdV1, err := crdClientSet.ApiextensionsV1().CustomResourceDefinitions().Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	crd := &apiextensions.CustomResourceDefinition{}
	err = apiextensionsv1.Convert_v1_CustomResourceDefinition_To_apiextensions_CustomResourceDefinition(crdV1, crd, nil)
	if err != nil {
		return err
	}

	validator, err := NewValidatorFromCRDs(*crd)
	if err != nil {
		return err
	}
	validators[name] = validator

	return nil
}
