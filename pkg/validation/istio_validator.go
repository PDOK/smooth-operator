package validation

/**
 * This logic is mostly copied from github.com/istio/istio/pkg/config/crd/validation
 */

import (
	"context"
	"fmt"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	apiextval "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/validation"
	structuralschema "k8s.io/apiextensions-apiserver/pkg/apiserver/schema"
	"k8s.io/apiextensions-apiserver/pkg/apiserver/schema/cel"
	structuraldefaulting "k8s.io/apiextensions-apiserver/pkg/apiserver/schema/defaulting"
	structurallisttype "k8s.io/apiextensions-apiserver/pkg/apiserver/schema/listtype"
	structuralpruning "k8s.io/apiextensions-apiserver/pkg/apiserver/schema/pruning"
	"k8s.io/apiextensions-apiserver/pkg/apiserver/validation"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	celconfig "k8s.io/apiserver/pkg/apis/cel"
	"sigs.k8s.io/yaml"
	"slices"
	"strings"
)

// Validator returns a new validator for custom resources
// Warning: this is meant for usage in tests only
type Validator struct {
	byGvk      map[schema.GroupVersionKind]validation.SchemaCreateValidator
	structural map[schema.GroupVersionKind]*structuralschema.Structural
	cel        map[schema.GroupVersionKind]*cel.Validator
	// If enabled, resources without a validator will be ignored. Otherwise, they will fail.
	SkipMissing bool
}

func (v *Validator) ValidateCustomResourceYAML(data string) error {
	obj := &unstructured.Unstructured{}
	if err := yaml.Unmarshal([]byte(data), obj); err != nil {
		return err
	}

	return v.ValidateCustomResource(obj)
}

func (v *Validator) ValidateCustomResource(o runtime.Object) error {
	content, err := runtime.DefaultUnstructuredConverter.ToUnstructured(o)
	if err != nil {
		return err
	}

	un := &unstructured.Unstructured{Object: content}
	vd, f := v.byGvk[un.GroupVersionKind()]
	if !f {
		if v.SkipMissing {
			return nil
		}
		return fmt.Errorf("failed to validate type %v: no validator found", un.GroupVersionKind())
	}
	// Fill in defaults
	structural := v.structural[un.GroupVersionKind()]
	structuraldefaulting.Default(un.Object, structural)
	if err := validation.ValidateCustomResource(nil, un.Object, vd).ToAggregate(); err != nil {
		return fmt.Errorf("%v/%v/%v: %v", un.GroupVersionKind().Kind, un.GetName(), un.GetNamespace(), err)
	}
	if err := structurallisttype.ValidateListSetsAndMaps(nil, structural, un.Object).ToAggregate(); err != nil {
		return fmt.Errorf("%v/%v/%v: %v", un.GroupVersionKind().Kind, un.GetName(), un.GetNamespace(), err)
	}
	pruneOpts := structuralschema.UnknownFieldPathOptions{TrackUnknownFieldPaths: true}
	unknownFieldPaths := structuralpruning.PruneWithOptions(un.DeepCopy().Object, structural, false, pruneOpts)
	unknownFieldPaths = FilterSliceInPlace(unknownFieldPaths, func(s string) bool {
		// Some CRDs don't spell out all the fields in metadata, and k8s doesn't care
		return !strings.HasPrefix(s, "metadata.")
	})
	if len(unknownFieldPaths) > 0 {
		return fmt.Errorf("%v/%v/%v: unknown fields %v", un.GroupVersionKind().Kind, un.GetName(), un.GetNamespace(), unknownFieldPaths)
	}

	errs, _ := v.cel[un.GroupVersionKind()].Validate(context.Background(), nil, structural, un.Object, nil, celconfig.RuntimeCELCostBudget)
	if errs.ToAggregate() != nil {
		return fmt.Errorf("%v/%v/%v: %v", un.GroupVersionKind().Kind, un.GetName(), un.GetNamespace(), errs.ToAggregate().Error())
	}
	return nil
}

func NewValidatorFromCRDs(crds ...apiextensions.CustomResourceDefinition) (*Validator, error) {
	v := &Validator{
		byGvk:      map[schema.GroupVersionKind]validation.SchemaCreateValidator{},
		structural: map[schema.GroupVersionKind]*structuralschema.Structural{},
		cel:        map[schema.GroupVersionKind]*cel.Validator{},
	}
	for _, crd := range crds {
		versions := crd.Spec.Versions
		if len(versions) == 0 {
			versions = []apiextensions.CustomResourceDefinitionVersion{{Name: crd.Spec.Version}} // nolint: staticcheck
		}
		//crd.Status.StoredVersions = slices.Map(versions, func(e apiextensions.CustomResourceDefinitionVersion) string {
		//	return e.Name
		//})
		errs := apiextval.ValidateCustomResourceDefinition(context.Background(), &crd)
		if len(errs) > 0 {
			return nil, fmt.Errorf("CRD %v is not valid: %v", crd.Name, errs.ToAggregate())
		}
		for _, ver := range versions {
			gvk := schema.GroupVersionKind{
				Group:   crd.Spec.Group,
				Version: ver.Name,
				Kind:    crd.Spec.Names.Kind,
			}
			crdSchema := ver.Schema
			if crdSchema == nil {
				crdSchema = crd.Spec.Validation
			}
			if crdSchema == nil {
				return nil, fmt.Errorf("crd did not have validation defined")
			}

			schemaValidator, _, err := validation.NewSchemaValidator(crdSchema.OpenAPIV3Schema)
			if err != nil {
				return nil, err
			}
			structural, err := structuralschema.NewStructural(crdSchema.OpenAPIV3Schema)
			if err != nil {
				return nil, err
			}

			v.byGvk[gvk] = schemaValidator
			v.structural[gvk] = structural
			// CEL programs are compiled and cached here
			if celv := cel.NewValidator(structural, true, celconfig.PerCallLimit); celv != nil {
				v.cel[gvk] = celv
			}

		}
	}

	return v, nil
}

// FilterInPlace retains all elements in []E that keep(E) returns true for.
// The array is *mutated in place* and returned.
// Use Filter to avoid mutation
func FilterSliceInPlace[E any](s []E, keep func(E) bool) []E {
	// find the first to filter index
	i := slices.IndexFunc(s, func(e E) bool {
		return !keep(e)
	})
	if i == -1 {
		return s
	}

	// don't start copying elements until we find one to filter
	for j := i + 1; j < len(s); j++ {
		if v := s[j]; keep(v) {
			s[i] = v
			i++
		}
	}

	clear(s[i:]) // zero/nil out the obsolete elements, for GC
	return s[:i]
}
