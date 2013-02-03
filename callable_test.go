package mogo

import (
	"testing"
)

func TestActDoesNotErrorWhenOneSetOfParametersAreAcceptable(t *testing.T) {
	const acceptableParams = "acceptable"
	c := callable{}
	c.IsCalled().WithParams("")
	c.IsCalled().WithParams(acceptableParams)

	c.actOn(acceptableParams)

	if c.err != nil {
		t.Errorf(`Expected callable to not have an error: %s`, c.err)
	}
}

func TestActDoesNotErrorWhenEmptyParametersAreAcceptable(t *testing.T) {
	c := callable{}
	c.IsCalled()
	c.IsCalled().WithParams("")

	c.actOn()

	if c.err != nil {
		t.Errorf(`Expected callable to not have an error: %s`, c.err)
	}
}
