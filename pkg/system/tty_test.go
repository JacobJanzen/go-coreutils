package system

import "testing"

func TestGetTTY(t *testing.T) {
	if _, err := GetTTY(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
