匿名函数：无函数名，也成为“闭包”，支持定义并**直接调用**；<u>一次性</u> or 赋给变量<u>多次</u>使用。

UC：

1. 代码封装，赋值变量 or 直接调用。
2. 闭包 ← 匿名函数 + 环境上下文。
3. 灵活简洁。
4. 并发，直接在匿名函数中定义并启动 goroutine。
5. 函数式编程，作参数传递 or 结果返回。

```go
// one-off
a := func(a, b int) int {
		return a + b
	}(1, 2)
```

```go
sum := func(a, b int) int {
		return a + b
	}
// multi-call
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





