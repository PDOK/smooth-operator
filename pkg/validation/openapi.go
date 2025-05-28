package validation

import (
	"context"
	"errors"
	"fmt"

	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/yaml"
)

var validators = map[string]*validator{}

// ValidateSchema validates incoming data against a OpenAPI specification
func ValidateSchema(yamlString string) error {
	data := &unstructured.Unstructured{}
	err := yaml.Unmarshal([]byte(yamlString), &data)
	if err != nil {
		return err
	}

	kind, ok := data.Object["kind"]
	if !ok {
		return errors.New("kind not found in yaml")
	}
	if val, ok := validators[kind.(string)]; ok {
		return val.ValidateCustomResourceYAML(yamlString)
	}
	return fmt.Errorf("no schemas found for CRD '%v'", kind)
}

func ApplySchemaDefaultsStr(yamlString string) (string, error) {
	data := &unstructured.Unstructured{}
	err := yaml.Unmarshal([]byte(yamlString), &data)
	if err != nil {
		return "", err
	}

	defaulted, err := ApplySchemaDefaults(data.Object)
	if err != nil {
		return yamlString, err
	}

	data.Object = defaulted
	result, err := data.MarshalJSON()
	if err != nil {
		return yamlString, err
	}

	return string(result), nil
}

func ApplySchemaDefaults(raw map[string]interface{}) (map[string]interface{}, error) {
	data := &unstructured.Unstructured{Object: raw}

	kind, ok := data.Object["kind"]
	if !ok {
		return raw, errors.New("kind not found in yaml")
	}
	if val, ok := validators[kind.(string)]; ok {
		err := val.ApplyDefaults(data)
		if err != nil {
			return raw, err
		}

		return data.Object, nil
	}
	return raw, fmt.Errorf("no schemas found for CRD '%v'", kind)
}

// AddSchema manually add an OpenAPI schema for a CRD
func AddValidator(schema apiextensionsv1.CustomResourceDefinition) error {
	if schema.Status.StoredVersions == nil {
		for _, v := range schema.Spec.Versions {
			if v.Storage {
				schema.Status.StoredVersions = append(schema.Status.StoredVersions, v.Name)
			}
		}
	}

	crd := &apiextensions.CustomResourceDefinition{}
	err := apiextensionsv1.Convert_v1_CustomResourceDefinition_To_apiextensions_CustomResourceDefinition(&schema, crd, nil)
	if err != nil {
		return err
	}

	val, err := newValidatorFromCRDs(*crd)
	if err != nil {
		return err
	}

	validators[schema.Spec.Names.Kind] = val
	return nil
}

// LoadSchemaForCRD extracts OpenAPI schemas for a specific CRD from a Kubernetes cluster
func LoadSchemasForCRD(cfg *rest.Config, namespace, name string) error {
	_ = namespace
	crdClientSet, err := clientset.NewForConfig(cfg)
	if err != nil {
		return err
	}

	crdV1, err := crdClientSet.ApiextensionsV1().CustomResourceDefinitions().Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	return AddValidator(*crdV1)
}
