匿名函数：无函数名，也成为“闭包”，支持定义并**直接调用**；One-off <u>一次性</u> or 赋给变量<u>多次</u>使用。

```go
a := func(a, b int) int {
		return a + b
	}(1, 2)
```

```go
sum := func(a, b int) int {
		return a + b
	}
sum(1, 2)
sum(3, 4)
```

全局匿名

```go
var (
	MyFunc = func(a, b int) int {
		return a + b
	}
)
a := MyFunc(1, 2)
```





