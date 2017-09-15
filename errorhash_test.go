package errors

import (
	"errors"
	"testing"
)

func TestAddErrorToErrorHash(t *testing.T) {
	e := NewErrorHash()
	e1 := errors.New("err 1")
	e2 := errors.New("err 2")

	e.Add("a", e1)
	e.Add("b", e2)
	e.Add("b", e1)

	if len(e) != 2 {
		t.Errorf("mismatch errors count, expected %d but got %d", 2, len(e))
	}

	if len(e["a"]) != 1 {
		t.Errorf("mismatch errors count for key a, expected %d but got %d", 1, len(e["a"]))
	}

	if len(e["b"]) != 2 {
		t.Errorf("mismatch errors count for key b, expected %d but got %d", 2, len(e["b"]))
	}
}

func TestErrorHashToString(t *testing.T) {
	e := NewErrorHash()
	e1 := errors.New("err 1")
	e2 := errors.New("err 2")

	e.Add("a", e1)
	if e.Error() != "a:[err 1]" {
		t.Errorf("expected message `%s` but got `%s`", "a:[err 1]", e.Error())
	}

	e.Add("b", e2)
	if e.Error() != "a:[err 1], b:[err 2]" {
		t.Errorf("expected message `%s` but got `%s`", "a:[err 1], b:[err 2]", e.Error())
	}

	e.Add("b", e1)
	if e.Error() != "a:[err 1], b:[err 2, err 1]" {
		t.Errorf("expected message `%s` but got `%s`", "a:[err 1], b:[err 2, err 1]", e.Error())
	}
}

func TestErrorHashHasString(t *testing.T) {
	e := NewErrorHash()
	e1 := errors.New("err 1")
	e2 := errors.New("err 2")
	e.Add("a", e1)
	e.Add("a", e2)

	if e.HasString("a", "err 2") != true {
		t.Error("expected true")
	}

	if e.HasString("b", "err 2") != false {
		t.Error("expected false")
	}
}

func TestErrorHashHasError(t *testing.T) {
	e := NewErrorHash()
	e1 := errors.New("err 1")
	e2 := errors.New("err 2")
	e.Add("a", e1)
	e.Add("a", e2)

	if e.HasError("a", e2) != true {
		t.Error("expected true")
	}

	if e.HasError("b", e2) != false {
		t.Error("expected false")
	}
}