只要类型实现了接口中的方法，那么该类型就属于该接口类型 = 实现了该接口。反过来说，一个接口类型的变量 `varInterface` 中可以包含任何类型的值，必须有一种方式来检测它的 **动态** 类型，即<u>运行时在变量中存储的值的实际类型</u>。

断言 Assertion 判断在某个时刻 `varInterface` 是否包含类型 T 的值。

```go
v := varInterface.(T)

if t, ok := varInterface.(T); ok {
    // do something
}

if _, ok := varInterface.(T); ok {
    // do something
}
```

接口变量的类型也可以使用一种特殊形式的 `type-switch` 来检测。

:warning: 此时不支持 `fallthrough`。

**UC：在处理来自于外部类型未知的数据时，比如解析 JSON or XML，类型测试和转换会非常有用。**

```go
switch v := varInterface.(type) {
case type1:
	// do something
case type2:
	// do something
default:
    // do something
}
```