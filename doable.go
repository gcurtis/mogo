package mogo

import (
	"reflect"
)

type doable struct {
	f interface{}
}

func (this *doable) AndDo(f interface{}) {
	this.f = f
}

func (this *doable) isDoable() bool {
	return this.f != nil
}

func (this *doable) run(params ...interface{}) R {
	v := reflect.ValueOf(this.f)

	in := make([]reflect.Value, len(params))
	for i, p := range params {
		if p == nil {
			// Reflect f's ith param to obtain a nil value for it.
			t := v.Type().In(i)
			in[i] = reflect.Zero(t)
		} else {
			in[i] = reflect.ValueOf(p)
		}
	}

	out := v.Call(in)
	ret := make([]interface{}, len(out))
	for i, o := range out {
		ret[i] = o.Interface()
	}

	return ret
}
