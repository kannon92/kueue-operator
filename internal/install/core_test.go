package install

import (
	"testing"
)

func TestBuildServiceAccount(t *testing.T) {
	got, err := BuildServiceAccount("../../assets/jobset/service_account.yaml")

	if err != nil {
		t.Errorf("error unexpected %s", err)
	}
	if got.Name != "jobset-controller-manager" {
		t.Errorf("got %v want %v", got.Name, "jobset-controller-manager")
	}
}

func TestBuildSecret(t *testing.T) {
	got, err := BuildSecrets("../../assets/jobset/secrets.yaml")

	if err != nil {
		t.Errorf("error unexpected %s", err)
	}
	if got.Name != "jobset-controller-manager" {
		t.Errorf("got %v want %v", got.Name, "jobset-controller-manager")
	}
}
