package system

import "testing"

func TestGetUsername(t *testing.T) {
	username, err := GetUsername(0)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if username != "root" {
		t.Errorf("expected root got %s", username)
	}
}
