package k8s

import (
	"context"
	"errors"
	"strings"

	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func GetConfigMap(k8sClient client.Client, namespace, prefix string, labels map[string]string) (*v1.ConfigMap, error) {
	if !strings.HasSuffix(prefix, "-") {
		prefix += "-"
	}

	configMapList := &v1.ConfigMapList{}
	err := k8sClient.List(context.TODO(), configMapList, client.InNamespace(namespace), client.MatchingLabels(labels))
	if err != nil {
		return nil, err
	}

	for _, configMap := range configMapList.Items {
		// If the name contains "-" after the prefix, it means that the prefix is longer
		// Example 'blobs' and 'blobs-premium'
		if strings.HasPrefix(configMap.Name, prefix) && !strings.Contains(configMap.Name[len(prefix):], "-") {
			return &configMap, nil
		}
	}

	return nil, errors.New("no configmap found with prefix " + prefix)
}

func GetSecret(k8sClient client.Client, namespace, prefix string, labels map[string]string) (*v1.Secret, error) {
	if !strings.HasSuffix(prefix, "-") {
		prefix += "-"
	}

	secretList := &v1.SecretList{}
	err := k8sClient.List(context.TODO(), secretList, client.InNamespace(namespace), client.MatchingLabels(labels))
	if err != nil {
		return nil, err
	}

	for _, secret := range secretList.Items {
		// If the name contains "-" after the prefix, it means that the prefix is longer
		// Example 'blobs' and 'blobs-premium'
		if strings.HasPrefix(secret.Name, prefix) && !strings.Contains(secret.Name[len(prefix):], "-") {
			return &secret, nil
		}
	}

	return nil, errors.New("no secret found with prefix " + prefix)
}
