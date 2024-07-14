### Type

数据类型**决定了变量所需的内存空间大小**，每个变量在内存中都需要分配一定的空间来存储其值。

| Type            | Desc                                                         |
| :-------------- | :----------------------------------------------------------- |
| **布尔型**：    | 布尔型的值只可以是常量 true 或者 false                       |
| **数字类型**：  | 整型 int 和浮点型 float32、float64。<br />Go 语言支持整型和浮点型数字，并且支持复数，其中位的运算采用补码。 |
| **字符串类型:** | 字符串就是一串定长的字符连接起来的字符序列。<br />Go 的字符串是由单个字节连接起来的。<br />Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本。 |
| **派生类型:**   | (a) 指针类型（Pointer）<br />(b) 数组类型<br />(c) 结构化类型(struct)<br />(d) Channel 类型<br />(e) 函数类型<br />(f) 切片类型<br />(g) 接口类型（interface）<br />(h) Map 类型 |

#### integer

Architecture-Specific Types 平台相关 (长度可变): int, uint, uintptr

| 序号       | Desc                                                         |
| :--------- | :----------------------------------------------------------- |
| **uint8**  | 无符号 8  位整型 (0 ~ 255)                                   |
| **uint16** | 无符号 16 位整型 (0 ~ 65535)                                 |
| **uint32** | 无符号 32 位整型 (0 ~ 4294967295)                            |
| **uint64** | 无符号 64 位整型 (0 ~ 18446744073709551615)                  |
| **int8**   | 有符号 8  位整型 (-128 ~ 127)                                |
| **int16**  | 有符号 16 位整型 (-32768 ~ 32767)                            |
| **int32**  | 有符号 32 位整型 (-2147483648 ~ 2147483647)                  |
| **int64**  | 有符号 64 位整型 (-9223372036854775808 ~ 9223372036854775807) |

**有符号 singed，可表示正负数和 0；无符号 unsinged 只可表示非负数。**

1. 对于有符号，最高位表示符号位：0 正数；1 负数。

```go
 1 // 0000 0001
-1 // 1000 0001
```

2. 正数源码/反码/补码一样 src = 1st's complement = 2nd complement。
3. 负数反码 = 源码符号位不变其他取反；补码 = 反码 + 1。

```go
 1  // 0000 0001  0000 0001  0000 0001
-1  // 1000 0001  1111 1110  1111 1111
```

4. **运算使用补码运算。**

:warning: Overflow

```go
var s int8 = 127 // int8: -128~127
s += 1           // -128
```

#### float(ing-point)

和平台无关，默认都是 0.0，但占内存空间 | 浮点数范围 | 精度也不同

| 序号           | Desc            |
| :------------- | :-------------- |
| **float32**    | 32 位浮点型数   |
| **float64**    | 64 位浮点型数   |
| **complex64**  | 32 位实数和虚数 |
| **complex128** | 64 位实数和虚数 |

#### byte

Golang里面没有专门的字符类型，如果要存储单个字符或者字母，一般使用 `byte` 来保存

#### bool

在 Go 语言中 bool 类型只能是**true or false**, 1 byte

### Conversion

Go 静态语言特性，不同类型之间复制需要**显式转换**，高位数类型 → 低位数会溢出 overflow (值不确定)

```go
var a int64 = 999999
var b int8 = int8(a) // overflow
```

**通常需要将基本数据类型转化为 string 类型，或者将 string 类型转化为基本数据类型。**

To-String:

1. `fmt.Sprintf` 返回格式化后的 str
2. `strconv.Iota` 将 int 转换成 str 
3. `strconv.Format*` 将 * 转换成 str

From-String:

1. `strconv.Parse*` 将 str 转换成 *

### Assertion

类型断言：将接口类型（表任意类型）的值转换为具体类型的操作。

```go
value, ok := x.(Type)
```

```go
var i interface{} = "Hello"

value, ok := i.(string)
if ok {
    fmt.Println("Value:", value)
} else {
    fmt.Println("Not a string type")
}
```

### Alias

**别名**：为现有类型定义一个新名称的方式，可以提高代码的**可读性和可维护性**。（带等号 `=`）

别名**不会创建新的类型**，它只是为现有类型提供了一个可选的替代名称。

```go
type NewType = ExistingType
```

```go
type Celsius float64
type PersonID string
type IntSlice []int
```

### Type Redefinition

**类型重定义**：将现有类型定义为不同的类型的操作。（**不带**等号 `=`）

类型重定义**会创建一个新的类型**，它与原始类型在类型系统中是**不同的**。所以必要时需要进行类型转换。

现有/原始类型也称为重定义类型的潜在类型；**潜在类型相同的不同类型的值之间是可以进行类型转换**。

```go
type ReDefType ExistingType
```

```go
type Celsius float64
type Fahrenheit float64

func main() {
    c := Celsius(25.0)
    f := Fahrenheit(c*9/5 + 32)
    
    fmt.Printf("%.2f degrees Celsius is %.2f degrees Fahrenheit\n", c, f)
}
```

