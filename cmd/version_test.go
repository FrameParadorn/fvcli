package cmd

import (
	"testing"
)

func TestCommandGetVersion(t *testing.T) {

	actual := getVersion()
	expected := "version v0.0.1"

	if actual != expected {
		t.Errorf("expected %s but got %s", expected, actual)
	}

}
