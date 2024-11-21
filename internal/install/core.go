package install

import (
	"fmt"
	"os"

	"k8s.io/client-go/kubernetes/scheme"

	corev1 "k8s.io/api/core/v1"
)

func BuildServiceAccount(serviceAccountPath string) (*corev1.ServiceAccount, error) {
	decode := scheme.Codecs.UniversalDeserializer().Decode
	stream, err := os.ReadFile(serviceAccountPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read %s got %w", serviceAccountPath, err)
	}
	obj, gKV, _ := decode(stream, nil, nil)
	if gKV.Kind == "ServiceAccount" {
		return obj.(*corev1.ServiceAccount), nil
	}
	return nil, fmt.Errorf("unable to decode service account")
}

func BuildSecrets(secretAccountPath string) (*corev1.Secret, error) {
	decode := scheme.Codecs.UniversalDeserializer().Decode
	stream, err := os.ReadFile(secretAccountPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read %s got %w", secretAccountPath, err)
	}
	obj, gKV, _ := decode(stream, nil, nil)
	if gKV.Kind == "Secret" {
		return obj.(*corev1.Secret), nil
	}
	return nil, fmt.Errorf("unable to decode secret")
}
