package mogo

import (
	"errors"
	"fmt"
	"reflect"
)

type Expectations struct {
	Returnable
	acceptableParams []interface{}
}

func (this *Expectations) act(args ...interface{}) error {
	if this.acceptableParams == nil {
		return nil
	}

	for i, a := range args {
		if !reflect.DeepEqual(this.acceptableParams[i], a) {
			return errors.New(fmt.Sprintf(`Arg %d did not match ("%v" != "%v").`, i, this.acceptableParams[i], a))
		}
	}

	return nil
}

func (this *Expectations) WithParams(args ...interface{}) *Returnable {
	if this.acceptableParams == nil {
		this.acceptableParams = make([]interface{}, len(args))
	}
	this.acceptableParams = args
	return &this.Returnable
}
