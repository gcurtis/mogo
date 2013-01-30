package mogo

import (
	"fmt"
)

type UnexpectedReturnError struct {
	method string
}

func (this *UnexpectedReturnError) Error() string {
	return fmt.Sprintf(`The mocked method "%s" was never expected to return anything.`, this.method)
}
