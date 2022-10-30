package system

import "testing"

func TestCheckProcCount(t *testing.T) {
	_, err := NProc()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCheckProcCountAll(t *testing.T) {
	_, err := NProcAll()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
