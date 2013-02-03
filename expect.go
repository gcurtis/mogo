package mogo

import (
	"reflect"
)

type retDo struct {
	returnable
	doable
}

type expect struct {
	retDo
	acceptableParams []interface{}
}

func (this *expect) act(args ...interface{}) (bool, R) {
	if this.acceptableParams == nil {
		return true, this.ret(args...)
	}

	for i, a := range args {
		if reflect.DeepEqual(this.acceptableParams[i], a) {
			return true, this.ret(args...)
		} else {
			return false, DefaultR
		}
	}

	return false, DefaultR
}

func (this *expect) ret(args ...interface{}) R {
	if this.isDoable() {
		return this.run(args...)
	}

	return this.returnable.ret()
}

func (this *expect) WithParams(args ...interface{}) *retDo {
	this.acceptableParams = args
	return &this.retDo
}
