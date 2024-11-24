package install

import (
	"fmt"
	"os"

	"k8s.io/client-go/kubernetes/scheme"

	corev1 "k8s.io/api/core/v1"
)

func ReadSecret(secretAccountPath, namespace string) (*corev1.Secret, error) {
	decode := scheme.Codecs.UniversalDeserializer().Decode
	stream, err := os.ReadFile(secretAccountPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read %s got %w", secretAccountPath, err)
	}
	obj, gKV, _ := decode(stream, nil, nil)
	if gKV.Kind == "Secret" {
		secret := obj.(*corev1.Secret)
		secret.ObjectMeta.Namespace = namespace
		return obj.(*corev1.Secret), nil
	}
	return nil, fmt.Errorf("unable to decode secret")
}

func ReadService(service, namespace string) (*corev1.Service, error) {
	decode := scheme.Codecs.UniversalDeserializer().Decode
	stream, err := os.ReadFile(service)
	if err != nil {
		return nil, fmt.Errorf("unable to read %s got %w", service, err)
	}
	obj, gKV, _ := decode(stream, nil, nil)
	if gKV.Kind == "Service" {
		service := obj.(*corev1.Service)
		service.ObjectMeta.Namespace = namespace
		return obj.(*corev1.Service), nil
	}
	return nil, fmt.Errorf("unable to decode service")
}
