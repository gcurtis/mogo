package mogo

import (
	"fmt"
)

type MockWriter struct {
	Mock
}

func (this *MockWriter) Write(p []byte) (n int, err error) {
	r := ActOnAndReturn(&this.Mock, "Write", p)
	n, _ = r[0].(int)
	err, _ = r[1].(error)

	return
}

func ExampleWriteNotCalled() {
	mock := MockWriter{}
	mock.Setup()
	mock.ExpectThat("Write").IsNotCalled()

	err := mock.Verify()
	if err == nil {
		fmt.Println("Mocks passed.")
	} else {
		fmt.Println(err)
	}

	// Output: Mocks passed.
}

func ExampleWriteCalledWithReturnValues() {
	mock := MockWriter{}
	mock.Setup()
	mock.ExpectThat("Write").IsCalled().AndReturns(R{10, nil})

	p := make([]byte, 10)
	n, err := mock.Write(p)
	fmt.Printf("Wrote %d bytes.", n)
	if err != nil {
		fmt.Println(err)
	}

	err = mock.Verify()
	if err != nil {
		fmt.Println(err)
	}

	// Output: Wrote 10 bytes.
}

func ExampleWriteWithParams() {
	mock := MockWriter{}
	mock.Setup()

	p := make([]byte, 10)
	mock.ExpectThat("Write").IsCalled().WithParams(p).AndReturns(R{10, nil})

	n, err := mock.Write(p)
	fmt.Printf("Wrote %d bytes.", n)
	if err != nil {
		fmt.Println(err)
	}

	err = mock.Verify()
	if err != nil {
		fmt.Println(err)
	}

	// Output: Wrote 10 bytes.
}
