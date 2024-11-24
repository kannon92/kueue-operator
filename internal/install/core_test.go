package install

import (
	"testing"
)

func TestReadSecret(t *testing.T) {
	got, err := ReadSecret("../../assets/jobset/secrets.yaml", "test")

	if err != nil {
		t.Errorf("error unexpected %s", err)
	}
	if got.Name != "jobset-controller-manager" {
		t.Errorf("got %v want %v", got.Name, "jobset-controller-manager")
	}
}
