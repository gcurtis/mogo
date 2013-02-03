package mogo

import (
	"testing"
)

func TestActReturnsTrueWhenAnyParameterIsAcceptable(t *testing.T) {
	e := expect{}
	acted, _ := e.act()

	if !acted {
		t.Error("Expected act to have returned true.")
	}
}

func TestActReturnsFalseWhenParametersAreNotAcceptable(t *testing.T) {
	e := expect{}
	e.WithParams("")

	acted, _ := e.act()

	if acted {
		t.Error("Expected act to have returned false.")
	}
}

func TestActReturnsTrueWhenParametersAreAcceptable(t *testing.T) {
	e := expect{}
	e.WithParams("")

	acted, _ := e.act("")

	if !acted {
		t.Error("Expected act to have returned true.")
	}
}

func TestWithParamsDoesNotReturnNil(t *testing.T) {
	e := expect{}

	actual := e.WithParams()
	if actual == nil {
		t.Errorf(`Expected WithParams to not have returned nil, instead got "%+v".`, actual)
	}
}

func TestThatDoableIsRun(t *testing.T) {
	e := expect{}

	doCalled := false
	e.AndDo(func() { doCalled = true })
	e.ret()

	if !doCalled {
		t.Error("Expected the doable to be called.")
	}
}
