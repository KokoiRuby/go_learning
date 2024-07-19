## Panic

**运行时**“恐慌”，会直接导致程序崩溃退出。

```bash
panic: runtime error: index out of range # panic desc & error val
goroutine 1 [running]:                   # goroutine id & status
main.main()                              # src, where panic is triggered
 /path/to/*.go:5 +0x3d
exit status 2                            # exit code
```

**panic 触发到程序中止的过程 = panicking？**

1. 创建 panic，控制权由 panic 触发代码行移交至所属函数。
2. 所属函数立即中止执行，并依次“冒泡”移交至上级。直到最外层 main，然后被 GC。
3. 程序崩溃，进程消失。

![image-20240716104315794](./panic.assets/image-20240716104315794.png)

**主动触发 panic 并指定值**

虽然 panic 可以接收空接口类型，但最好传入 error 类型或任何可序列化。

```go
func myFunc() {
	panic(errors.New("something went wrong"))
}
```

```go
type CustomError struct {
	Message string
	Code    int
}

func (e CustomError) Error() string {
	return fmt.Sprintf("code %d: %s", e.Code, e.Message)
}
myFunc
func myFunc() {
	panic(CustomError{Message: "custom error occurred", Code: 123})
}
```

## Recover

recover 用于**恢复** panic，无需任何参数，返回一个空接口类型，值为 panic 包含的值。

结合 defer + 匿名函数，是为了确保 recover 在函数移交控制权中止前一定能够捕捉。

**使用多条 defer 的时候要注意执行顺序。**

```go
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	myFunc()
}

func myFunc() {
	panic("something went wrong")
}
```

defer 还可以在引发新的 panic

```go
func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Recovered from panic: %v\n", err)
			// new panic
			panic("New panic in defer")
		}
	}()

	// orig panic
	panic("Original panic")
}
```

