package mogo

import (
	"reflect"
)

type Kind uint

const (
	Invalid Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Ptr
	Slice
	String
	Struct
	UnsafePointer
)

type retDo struct {
	returnable
	doable
}

type expect struct {
	retDo
	acceptableParams []interface{}
}

type any struct {
	i interface{}
}

func Any(i interface{}) any {
	return any{i: i}
}

func AnyKindOf(k Kind) any {
	return any{i: k}
}

func (this any) matches(i interface{}) bool {
	if k, ok := this.i.(Kind); ok {
		return reflect.TypeOf(i).Kind() == reflect.Kind(k)
	}

	t := reflect.TypeOf(i)
	anyT := reflect.TypeOf(this.i)
	return t == anyT
}

func (this *expect) act(args ...interface{}) (bool, R) {
	if this.acceptableParams == nil {
		return true, this.ret(args...)
	}

	for i, a := range args {
		p := this.acceptableParams[i]
		if q, ok := p.(any); ok {
			if q.matches(a) {
				return true, this.ret(args...)
			} else {
				return false, DefaultR
			}
		} else if reflect.DeepEqual(p, a) {
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
