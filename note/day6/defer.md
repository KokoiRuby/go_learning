defer：一种类栈的延时机制，通常用于资源的释放。

- Go 逐行编译执行到 defer 的时候，会将其语句入栈，暂时不执行
- 在函数 return 之前会依次出栈执行

```go
// file
defer file Close()

// db conn
defer connect.close()
```