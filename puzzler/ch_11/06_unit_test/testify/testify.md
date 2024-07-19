Go standard test package `testing` does a decent job in testing but it still lacks some useful functionality like **assertion**, **mocking**, and test **suite**.

[Testify](https://github.com/stretchr/testify) as ext. to fulfill.

```bash
# install
$ go get github.com/stretchr/testify
```

### [Assertion](https://github.com/stretchr/testify?tab=readme-ov-file#assert-package)

An assertion is a statement in a program that checks a certain condition or expression and **ensures that it evaluates to true**.

- `[assert|require].Equal(t, expected, actual, "msg")`
- `[assert|require].NotEqual(t, not_expected, actual, "msg")`
- `[assert|require].Nil(t, interface{}, "msg")`
- `[assert|require].NotNil(t, interface{}, "msg")`

`assert` will log the failure and continue with the test, while `require` will log the failure and immediately stop the test.

### [Mock](https://github.com/stretchr/testify?tab=readme-ov-file#mock-package)

A mock is a **simulated** object as **stub** that mimics the behavior of real objects in controlled ways. This allows for testing the functionality of the code **without relying on** the actual implementations of its dep.

- `mockObj.On(MethodName, args...interface{}).Return(ret...interface{})`

```go
// a mock obj that mock Printer interface
type PrinterMock struct {
	mock.Mock
}

func (mk *PrinterMock) Print(paperSize string, content string) error {
	// stub it out instead of do
    args := mk.Called(paperSize, content)
	fmt.Println("Printing from mock objct")
	return args.Error(0)
}
```

```go
// create mock
printer := &PrinterMock{}
// set expectation that the mocked function will receive
printer.On("Print", "A4", "Hello World").Return(nil)
```

[Mockery](https://github.com/vektra/mockery) is a very handy package that generates mocks for you.

```bash
$ go get github.com/vektra/mockery/v2/.../
$ go rungithub.com/vektra/mockery/v2 --all
```

### [Suite](https://github.com/stretchr/testify?tab=readme-ov-file#suite-package)

When there’re way too many tests in your software, it could be cumbersome to manage them. **A set of test cases** is called a test suite → **categorize**

It provides you high **flexibility** about **how and when** you execute your code in a test suite. → **setup → before/after tc → cleanup**

- `SetupSuite()`
- `TearDownSuite()`
- `SetupTest()`
- `TearDownTest()`

```go
type CalculatorTestSuite struct {
	suite.Suite
	StartingNumber int
}

func TestCalculatorTestSuite(t *testing.T) {
	suite.Run(t, new(CalculatorTestSuite))
}
```