package install

import (
	"fmt"
	"os"

	v1 "github.com/kannon92/kueue-operator/api/v1"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

func BuildJobSetDeployment(jobSetSpec *v1.JobSetSpec) (*appsv1.Deployment, error) {
	jobSetDeployment := "../../assets/jobset/deployment.yaml"
	decode := scheme.Codecs.UniversalDeserializer().Decode
	stream, err := os.ReadFile(jobSetDeployment)
	if err != nil {
		return nil, fmt.Errorf("unable to read %s got %w", jobSetDeployment, err)
	}
	obj, gKV, _ := decode(stream, nil, nil)
	if gKV.Kind != "Deployment" {
		return obj.(*appsv1.Deployment), nil
	}

	deployment := obj.(*appsv1.Deployment)
	for i, val := range deployment.Spec.Template.Spec.Containers {
		if val.Name == "manager" {
			deployment.Spec.Template.Spec.Containers[i].Image = jobSetSpec.JobSetImage
		}
		if val.Name == "kube-rbac-proxy" {
			deployment.Spec.Template.Spec.Containers[i].Image = jobSetSpec.Proxy
		}
	}
	return deployment, nil
}
