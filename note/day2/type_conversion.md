Go 静态语言特性，不同类型之间复制需要**显式转换**，高位数类型 → 地位数会溢出 overflow (值不确定)

```go
var a int64 = 999999
var b int8 = int8(a) // overflow
```

**经常需要将基本数据类型转化为 string 类型，或者将 string 类型转化为基本数据类型**

To-String:

1. `fmt.Sprintf` 返回格式化后的 str
2. `strconv.Ita` 将 int 转换成 str
3. `strconv.Format*` 将 * 转换成 str

From-String:

1. `strconv.Parse*` 将 str 转换成 *

### [fmt](https://pkg.go.dev/fmt) ([CN](https://www.topgoer.com/%E5%B8%B8%E7%94%A8%E6%A0%87%E5%87%86%E5%BA%93/fmt.html))

%* as placeholder. width (space padding) & precision. 

键盘输入：`fmt.Scanln()`, `fmt.Scanf()`
