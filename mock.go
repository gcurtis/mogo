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
		if call.expectations.f != nil {
			return call.expectations.run(args...)
		}
		return call.expectations.returns
	}

	return make(R, 32)
}

func ActOnAndReturnOne(mock *Mock, method string, args ...interface{}) interface{} {
	call, ok := mock.calls[method]
	if ok && call.shouldBeCalled {
		mock.err = call.actOn(args...)
		if call.expectations.f != nil {
			return call.expectations.run(args...)[0]
		}
		return call.expectations.returns[0]
	}

	return make(R, 32)
}

func NewMock() *Mock {
	return &Mock{
		calls: make(map[string]*callable),
	}
}

type Mock struct {
	calls map[string]*callable
	err   error
}

func (this *Mock) Setup() {
	this.calls = make(map[string]*callable)
}

func (this *Mock) ExpectThat(method string) *callable {
	call := &callable{
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
