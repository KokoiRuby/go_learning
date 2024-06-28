**字符串：一串字符连接起来的字符定长序列，go 语言的字符串是由单个字节连接起来的。**

- **不可变**：一旦声明就不可二次赋值，提高<u>并发安全性</u>和存储利用率 + 只需要分配一块内存空间
- O(1) 获取字符串长度 len

```go
var s1 string = "this is a string"
s2 := `
this
is
yet
another
string.
`
```

#### essense

string 是一个“描述符”，它本身并不真正存储字符串数据，**二元组**：一个指向底层存储的**指针 ptr**和字符串的**长度 len**字段组成的

指针指向的是一个字节数组 `[x]byte`

传参时，传递的是 FD，不会涉及底层数据拷贝，内存里就只有一份内存空间。

```go
type StringHeader struct {
	Data uintptr
	Len  int
}

// len(5) + Data → [h][e][l][l][o] = [5]byte{"h", "e", "l", "l", "o"}
var s string = "hello"
```

如果真的要修改，**可转化为 byte 切片或者 rune 切片，再转化为字符串**

```go
str1 := "big"
b := []byte(str1)
b[0] = 'p'
fmt.Println(string(b))

str2 := "pig"
r := []rune(str2)
r[0] = 'b'
fmt.Println(string(r))
```

#### vss

- `string` 是一系列 Unicode 字符的集合，用于表示文本数据。字符串的每个字符都由一个或多个字节表示，取决于字符的 Unicode 编码。
- `rune/int32` 是 Go 语言中用于表示 Unicode 字符的类型，每个 `rune` 表示一个 Unicode 码点，可以表示任何 Unicode 字符。
- `byte/uint8` 是 Go 语言中用于表示单个字节的类型，当处理 ASCII 字符或字节数据时，可以使用 `byte` 类型。



<img src="https://miro.medium.com/v2/resize:fit:1050/1*b3TZICZOHODu0gWJdmH2KA.png" alt="img" style="zoom:67%;" />



<img src="https://miro.medium.com/v2/resize:fit:1050/1*BxXZA-6Xr43TP8r0xjyn7A.png" alt="img" style="zoom:50%;" />

#### utf-8

一种编码格式，决定 Unicode 码点在计算机中如何存储和表示，即 Unicdode char 映射 → Byte sequence.

一个 Unicode 码点，根据语言不同，一般 1~4 Bytes 不等。