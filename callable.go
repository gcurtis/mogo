package mogo

import (
	"bytes"
	"fmt"
)

const expectFmt = `"%s" wasn't called with the correct arguments.

Actual:
%s
%s`

type callable struct {
	method     string
	expectCall bool
	called     bool
	set        bool
	expects    []*expect
	err        error
}

func (this *callable) actOn(args ...interface{}) R {
	if !this.expectCall {
		this.err = fmt.Errorf(`"%s" was called but should not have been.`, this.method)
		return DefaultR
	}

	for _, e := range this.expects {
		acted, r := e.act(args...)
		if acted {
			this.called = true
			return r
		}
	}

	var actBuf bytes.Buffer
	for i, a := range args {
		actBuf.WriteString(fmt.Sprintf("\tParam %d: %#v\n", i, a))
	}

	var expBuf bytes.Buffer
	for i, e := range this.expects {
		expBuf.WriteString(fmt.Sprintf("\nExpected %d:\n", i))
		for i, ep := range e.acceptableParams {
			expBuf.WriteString(fmt.Sprintf("\tParam %d: %#v\n", i, ep))
		}
	}

	// Remove trailing newlines
	actStr := actBuf.String()
	actStr = actStr[:len(actStr)-1]
	expStr := expBuf.String()
	expStr = expStr[:len(expStr)-1]

	this.err = fmt.Errorf(expectFmt, this.method, actStr, expStr)
	return DefaultR
}

func (this *callable) verify() error {
	if this.err != nil {
		return this.err
	}

	if this.expectCall && !this.called {
		return fmt.Errorf("\"%s\" should have been called but was not.\n\r", this.method)
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
	} else {
		this.expectCall = true
		this.set = true
	}

	e := &expect{}
	this.expects = append(this.expects, e)
	return e
}
