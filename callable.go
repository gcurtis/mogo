package mogo

import (
	"errors"
	"fmt"
)

type callable struct {
	method         string
	shouldBeCalled bool
	wasCalled      bool
	expectations   expectations
}

func (this *callable) actOn(args ...interface{}) error {
	this.wasCalled = true
	if this.shouldBeCalled {
		return this.expectations.act(args...)
	}

	return errors.New(fmt.Sprintf(`%s was called but should not have been.`, this.method))
}

func (this *callable) verify() error {
	if this.shouldBeCalled && !this.wasCalled {
		return errors.New(fmt.Sprintf(`%s should have been called but was not.`, this.method))
	}

	return nil
}

func (this *callable) IsNotCalled() {
	this.shouldBeCalled = false
}

func (this *callable) IsCalled() *expectations {
	this.shouldBeCalled = true
	return &this.expectations
}
