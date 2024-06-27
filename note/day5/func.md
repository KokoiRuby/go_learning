**函数**是基本的代码块，用于执行一个任务。实现封装，使用调用。

至少包含一个名为 main 的函数，作为该程序的入口。

```go
func func_name ( [param list] ) [return_types] {
   // body
}
```

- `func` 函数声明关键字
- `func_name` 函数名
- `param list` 形参列表 + `return_types` 返回值，可多返回（使用 `_` 忽略返回） → 函数签名 Signature