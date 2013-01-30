package mogo

func ActOn(mock *Mock, method string, args ...interface{}) {
	call, ok := mock.calls[method]
	if ok {
		mock.err = call.actOn(args...)
	}
}

func ActOnAndReturn(mock *Mock, method string, args ...interface{}) R {
	call, ok := mock.calls[method]
	if ok && call.shouldBeCalled {
		mock.err = call.actOn(args...)
		return call.expectations.returns
	}

	return make(R, 32)
}

func ActOnAndReturnOne(mock *Mock, method string, args ...interface{}) interface{} {
	call, ok := mock.calls[method]
	if ok && call.shouldBeCalled {
		mock.err = call.actOn(args...)
		return call.expectations.returns[0]
	}

	return make(R, 32)
}

func NewMock() *Mock {
	return &Mock{
		calls: make(map[string]*Callable),
	}
}

type Mock struct {
	calls map[string]*Callable
	err   error
}

func (this *Mock) Setup() {
	this.calls = make(map[string]*Callable)
}

func (this *Mock) ExpectThat(method string) *Callable {
	call := &Callable{
		method: method,
	}

	this.calls[method] = call
	return call
}

func (this *Mock) Verify() error {
	if this.err != nil {
		return this.err
	}

	for _, call := range this.calls {
		err := call.verify()
		if err != nil {
			return err
		}
	}

	return nil
}
