package utils

import (
	"testing"
)

// TestGetVersion tests
func TestGetVersion(t *testing.T) {
	out := GetVersion()
	if out == "" || len(out) == 0 {
		t.Errorf("Want strings but got nil")
	}
}
