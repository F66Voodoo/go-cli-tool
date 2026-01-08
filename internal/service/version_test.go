package service

import "testing"

func TestVersion_NotEMpty(t *testing.T) {
	if Version() == "" {
		t.Fatal("Version() return empty string")
	}
}
