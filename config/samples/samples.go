package samples

import (
	_ "embed"
	v1 "github.com/pdok/smooth-operator/api/v1"
	"sigs.k8s.io/yaml"
)

//go:embed v1_ownerinfo.yaml
var ownerInfoContent string

func OwnerInfoSample() (*v1.OwnerInfo, error) {
	var sample v1.OwnerInfo
	err := yaml.Unmarshal([]byte(ownerInfoContent), &sample)
	return &sample, err
}
