package system

import (
	"os"
	"testing"
)

func TestGetUsername(t *testing.T) {
	username, err := GetUsername(0)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if username != "root" {
		t.Errorf("expected root got %s", username)
	}
}

func TestGetLogin(t *testing.T) {
	username, err := GetLogin()

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if username != os.Getenv("USER") {
		t.Errorf("expected %s got %s", os.Getenv("USER"), username)
	}
}
