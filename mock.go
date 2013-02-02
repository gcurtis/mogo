package mogo

func act(mock *Mock, method string, args ...interface{}) R {
	call, ok := mock.calls[method]
	if ok {
		call.actOn(args...)
		if call.expect.isDoable() {
			return call.expect.run(args...)
		}

		return call.expect.returns
	}

	return make(R, 32)
}

func ActOn(mock *Mock, method string, args ...interface{}) {
	act(mock, method, args...)
}

func ActOnAndReturn(mock *Mock, method string, args ...interface{}) R {
	return act(mock, method, args...)
}

func ActOnAndReturnOne(mock *Mock, method string, args ...interface{}) interface{} {
	return act(mock, method, args...)[0]
}

type verifiable interface {
	Verify() error
}

func VerifyAll(mocks ...verifiable) (err error) {
	for _, m := range mocks {
		err = m.Verify()
		if err != nil {
			return
		}
	}

	return
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

func (this *Mock) Setup() *Mock {
	this.calls = make(map[string]*callable)
	return this
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
