mogo
====

**WARNING: MOGO IS CURRENTLY PRE-RELEASE.**

The goal of mogo is to provide a mocking framework for Go that doesn't rely on code generation. Since Go doesn't have a way creating new types at runtime, mogo works by providing helper functions that allow you to create mocks using just a few lines of code.

Quickstart
----------

To begin creating a mock, we need to implement an interface. This example uses the `io.Writer` interface:

```go
type MockWriter struct{}

func (this *MockWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}
```

So far we've just stubbed out the `io.Writer` interface. Now lets add some mogo methods to make it a true mock:

```go
type MockWriter struct {
	// Embedding the Mock struct automatically adds mocking methods
	mogo.Mock
}

func (this *MockWriter) Write(p []byte) (n int, err error) {
	// This method allows mogo to verify your mocks. It takes a pointer to our
	// embedded Mock struct, the name of the method we're mocking, and any
	// parameters the method takes.
	r := mogo.ActOnAndReturn(&this.Mock, "Write", p)
	
	// ActOnAndReturn returns an array of interfaces. Each item in this array
	// corresponds to a value we should return. Notice we use the multiple
	// return form for the type assertions so we don't panic if we need to
	// return nil.
	n, _ = r[0].(int)
	err, _ = r[1].(error)
	return
}
```

We're done creating our mock! Notice that without comments our mock is less than 10 lines of code. We can now use it in our tests like so:

```go
func TestSomethingWithWriter(t *testing.T) {
	mockWriter := MockWriter{}
	mockWriter.Setup()
	mockWriter.ExpectThat("Write").IsCalled().AndReturn(10, nil)
	
	// Do some test stuff...
	
	err := mockWriter.Verify()
	if err != nil {
		t.Errorf("Mock verification failed: %v", err)
	}
}
```

To see more examples on how to use mogo, take a look at [examples_test.go](examples_test.go).
