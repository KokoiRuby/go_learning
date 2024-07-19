### error

error 类型其实是一个接口类型，不接受任何参数，返回错误信息的**字符串**表示形式。

:construction_worker: 函数声明的返回列表中，声明一个 error 类型返回；调用函数时，判 nil 决定是否进入错误处理。

:construction_worker: **卫述语句 Guard Clause**：防御性编程风格，常用于检查/处理函数/方法开头处的特殊情况或错误条件。目的是尽早返回错误。

```go
type error interface {
    Error() string
}

func myFunc(...) (... err error) {
	//
    err = errors.New("empty request")
    //
}
_, err := myFunc(...)
if err != nil {
    // err handling
}
// continue
```

**定义错误**

```go
err := errors.New("err string")
```

**结构化错误**

```go
type PathError struct {
    Op string    // operation
    Path string  // path/to/file
    Err error    // Returned by the system call.
}

func (e *PathError) String() string {
    return e.Op + " " + e.Path + ": "+ e.Err.Error()
}
```

`fmt.Errorf` **格式化**的错误值，返回一个格式化的字符串。

```go
err := fmt.Errorf("an error occurred: %s", "details about the error")
fmt.Println(err)
```

∵ error 是一个接口类型，即使同为 error 类型的错误值，**实际类型也可能不同**。**如何判断呢？**

1. 对于**不同类型**，一般使用**类型断言或类型 switch 语句**来判断。

```go
func underlyingError(err error) error {
    switch err := err.(type) {
    case *os.PathError:
    	return err.Err
    case *os.LinkError:
    	return err.Err
    case *os.SyscallError:
    	return err.Err
    case *exec.Error:
    	return err.Err
    }
    return err
}

if e, ok := err.(type); ok {
    // continue
}
```

2. 对于**类型相同**的，一般直接使用**if + 判等或 switch**操作来判断。

```go
printError := func(i int, err error) {
    if err == nil {
    	fmt.Println("nil error")
   	 return
    }
    err = underlyingError(err)
    switch err {
    case os.ErrClosed:
    	fmt.Printf("error(closed)[%d]: %s\n", i, err)
    case os.ErrInvalid:
    	fmt.Printf("error(invalid)[%d]: %s\n", i, err)
    case os.ErrPermission:
    	fmt.Printf("error(permission)[%d]: %s\n", i, err)
    }
    
    // if err == os.ErrClosed { ... }
    // if err == os.ErrInvalid { ... }
    // if err == os.ErrPermission { ... }
}

```

3. 对于**类型未知**，只能使用其**错误信息的字符串**表示形式来做判断。

```go
func doSomething() error {
	return errors.New("file does not exist")
}

err := doSomething()
if err != nil {
    if err.Error() == "file does not exist" {
		fmt.Println("Error: The file does not exist.")
	} else if err.Error() == "permission denied" {
		fmt.Println("Error: Permission denied.")
	} else {
		fmt.Println("Error:", err)
	}
}
```

**如何根据实际情况给出恰当的错误值？**

1. 参考内置 error **树形架构**，root & leaves，用统一字段建立起可追根溯源的**链式**错误关联。

```go
error
├── *os.PathError
├── *os.LinkError
├── *os.SyscallError
├── os.ErrInvalid
├── os.ErrPermission
├── os.ErrExist
├── os.ErrNotExist
├── os.ErrClosed
├── *net.OpError
├── *net.DNSError
├── *net.AddrError
├── *net.InvalidAddrError
├── *io.EOF
├── *io.ErrUnexpectedEOF
├── *io.ErrNoProgress
├── UserDefinedError
│   ├── errors.New
│   ├── fmt.Errorf
│   └── CustomErrorStruct
```

2. **扁平化** Flatting，简单；建议小写**私有**，然后创建**公有获取错误值以及判等**的函数。

```go
var (
	errInvalidInput   = errors.New("invalid input")
	errNotFound       = errors.New("not found")
)

func ErrInvalidInput() error {
	return errInvalidInput
}

func ErrNotFound() error {
	return errNotFound
}

func IsInvalidInput(err error) bool {
	return err == errInvalidInput
}

func IsNotFound(err error) bool {
	return err == errNotFound
}
```