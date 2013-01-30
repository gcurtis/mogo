package mogo

import (
	"errors"
	"fmt"
)

type Callable struct {
	method         string
	shouldBeCalled bool
	wasCalled      bool
	expectations   Expectations
}

func (this *Callable) actOn(args ...interface{}) error {
	this.wasCalled = true
	if this.shouldBeCalled {
		return this.expectations.act(args...)
	}

	return errors.New(fmt.Sprintf(`%s was called but should not have been.`, this.method))
}

func (this *Callable) verify() error {
	if this.shouldBeCalled && !this.wasCalled {
		return errors.New(fmt.Sprintf(`%s should have been called but was not.`, this.method))
	}

	return nil
}

func (this *Callable) IsNotCalled() {
	this.shouldBeCalled = false
}

func (this *Callable) IsCalled() *Expectations {
	this.shouldBeCalled = true
	return &this.expectations
}
