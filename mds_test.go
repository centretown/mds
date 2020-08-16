package main

import (
	"testing"
)

// TestSetFlags -
func TestMakeSettings(t *testing.T) {
	settings, err := MakeSettings()
	if settings == nil {
		t.Fatalf("%v", "nil settings")
	}
	if err != nil {
		t.Logf("%v", err)
	}
	t.Logf("%+v", settings)
}
