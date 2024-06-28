`panic` 接受 `interface{} `类型的值作为参数，也可以接收 `error` 类型变量，输出错误信息，并退出程序。

```go
func panic(v interface{})
```

:construction_worker: Practice：抛出一个 `panic` 异常，在 `defer` 中 (匿名) 通过 `recover` 捕获并处理。

```go
func fn1(a int, b int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
    }()           // don't forget () to call anonymous
	return a / b
}
```

