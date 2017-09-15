package errors

import (
	"errors"
	"testing"
)

func TestAddErrorToErrors(t *testing.T) {
	var e Errors
	e1 := errors.New("err 1")
	e2 := errors.New("err 2")

	e.Add(e1)
	e.Add(e2)

	if len(e) != 2 {
		t.Errorf("mismatch errors count, expected %d but got %d", 2, len(e))
	}
}

func TestErrorsToString(t *testing.T) {
	var e Errors
	e1 := errors.New("err 1")
	e2 := errors.New("err 2")

	e.Add(e1)
	if e.Error() != "err 1" {
		t.Errorf("expected message `%s` but got `%s`", "err 1", e.Error())
	}

	e.Add(e2)
	if e.Error() != "err 1, err 2" {
		t.Errorf("expected message `%s` but got `%s`", "err 1, err 2", e.Error())
	}
}
