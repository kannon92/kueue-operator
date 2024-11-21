package install

import (
	"testing"

	v1 "github.com/kannon92/kueue-operator/api/v1"
)

func TestBuildJobSetDeployment(t *testing.T) {

	got, err := BuildJobSetDeployment(&v1.JobSetSpec{JobSetImage: "testMe", Proxy: "proxyTest"})

	if err != nil {
		t.Errorf("error unexpected %s", err)
	}
	for _, val := range got.Spec.Template.Spec.Containers {
		if val.Name == "manager" {
			if val.Image != "testMe" {
				t.Error("manager image supposed to be equal to testMe")
			}
		}
		if val.Name == "kube-rbac-proxy" {
			if val.Image != "proxyTest" {
				t.Error("kube rbac proxy supposed to be equal to proxyTest")
			}
		}
	}

}
