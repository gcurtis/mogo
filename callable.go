package mogo

import (
	"fmt"
)

type callable struct {
	method     string
	expectCall bool
	called     bool
	set        bool
	expect     expect
	err        error
}

func (this *callable) actOn(args ...interface{}) {
	this.called = true
	if !this.expectCall {
		this.err = fmt.Errorf(`%s was called but should not have been.`, this.method)
		return
	}

	this.expect.act(args...)
}

func (this *callable) verify() error {
	if this.err != nil {
		return this.err
	}

	if this.expectCall && !this.called {
		return fmt.Errorf(`%s should have been called but was not.`, this.method)
	}

	return nil
}

func (this *callable) IsNotCalled() {
	if this.set && this.expectCall {
		this.err = fmt.Errorf(`"%s" was expected to be called and not called.`, this.method)
		return
	}

	this.expectCall = false
	this.set = true
}

func (this *callable) IsCalled() *expect {
	if this.set && !this.expectCall {
		this.err = fmt.Errorf(`"%s" was expected to be called and not called.`, this.method)
		return &this.expect
	}

	this.expectCall = true
	this.set = true
	return &this.expect
}
