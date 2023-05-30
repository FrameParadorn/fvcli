package new

import (
	"testing"
)

func TestCreateCommandShouldReturnUseEqualNew(t *testing.T) {

	new := NewNew()
	command := new.CreateCommand()

	got := command.Use
	expected := "new"

	if got != expected {
		t.Errorf("CreateCommand 'new' fail got %s but want %s", got, expected)
	}

}

func TestRunFunctionNewCommandShouldWithoutError(t *testing.T) {

	new := NewNew()
	command := new.CreateCommand()

	err := command.RunE(nil, nil)
	if err != nil {
		t.Errorf("Run function fail with return error: %e", err)
	}

}
