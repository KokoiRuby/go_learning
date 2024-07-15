### String

**字符串：一串字符连接起来的字符定长序列，go 语言的字符串是由单个字节连接起来的。**

- **不可变**：一旦声明就不可二次赋值，提高<u>并发安全性</u>和存储利用率 + 只需要分配一块内存空间
- O(1) 获取字符串长度 len

```go
var s1 string = "this is a string"
s2 := `
this
is
yetanother
string.
`
```

#### escape

| Char |  Desc   |
| :--: | :-----: |
| `\t` |   tab   |
| `\n` | newline |
| `\\` |    \    |
| `\r` |  enter  |
| `"`  |    "    |

#### essense

string 是一个“描述符”，它本身并不真正存储字符串数据，**二元组**：一个指向底层存储的**指针 *Data**和字符串的**长度 len**字段组成的

指针指向的是一个字节数组 `[x]byte`

![Go 语言字符串和字节数组的实现原理| Go 语言设计与实现](https://img.draveness.me/2019-12-31-15777265631608-in-memory-string.png)

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

### fmt

| 函数       | 主要用途                                         | 输出行为                               | 返回值               |
| ---------- | ------------------------------------------------ | -------------------------------------- | -------------------- |
| `Println`  | 打印一系列值，自动添加空格和换行符               | 每个值之间有空格，末尾有换行符         | 无返回值             |
| `Printf`   | 格式化输出，根据格式化动词指定输出格式           | 按照格式化动词指定的格式输出，无换行符 | 无返回值             |
| `Sprintln` | 打印一系列值，自动添加空格和换行符，但返回字符串 | 每个值之间有空格，末尾有换行符         | 返回格式化后的字符串 |

```go
name := "Alice"
age := 30
// Name: Alice Age: 30
fmt.Println("Name:", name, "Age:", age)

// Name: Alice Age: 30
fmt.Printf("Name: %s Age: %d\n", name, age)

result := fmt.Sprintln("Name:", name, "Age:", age)
// Name: Alice Age: 30
fmt.Print(result)
```

### Encoding

Go 的所有源码都按照 Unicode 编码规范中 UTF-8 编码格式进行编码。如果出现非 UTF-8 字符则报错。

Go 支持将一个整数转换成一个 string 类型的值，但前提是整数值可代表一个有效的 unicode 码点。否则将表示为一个高亮的 `?` 字符。

#### ASCII

American Standard Code for Information Interchange 美国信息交换标准代码。

后定位 ISO 646 标准，**仅适用于所有拉丁字母**。

ASCII 使用单个字节（byte）的二进制数来编码一个字符。

#### Unicode

更通用，世界上**所有自然语言中的每个字符**都有一个唯一的二进制编码。

计算机内部，**所有的字符都会被编码为整数**，每个特定整数都称为一个代码点。超出范围的表示为一个高亮的 `?`。

使用十六进制表示：U+0061 → 'a'。

支持 3 种**编码格式**：UTF-8/16/32

#### UTF-8

UTF-8 可变宽，会用一个或多个字节的二进制数来表示某个字符，最多使用四个字节。即**字符和字节转换**。

一个受支持的字符总是可以由 UTF-8 编码为一个**字节序列**。

### vs

一个 string 类型的值可被拆分为一个**字符序列 []rune** ，也可被拆分为一个**字节序列 []byte**。

**一个 string 类型的值在底层就是一个能够表达若干个 UTF-8 编码值的字节序列。**

rune 是 Go 语言特有的一个基本数据类型，代表一个 Unicode 字符。

```go
type rune = int32 // alias
```

<img src="https://miro.medium.com/v2/resize:fit:1050/1*b3TZICZOHODu0gWJdmH2KA.png" alt="img" style="zoom:67%;" />

<img src="https://miro.medium.com/v2/resize:fit:1050/1*BxXZA-6Xr43TP8r0xjyn7A.png" alt="img" style="zoom:50%;" />

### [strings](https://pkg.go.dev/strings)

:cry: ​`string` 

- 不可变；如果要修改就只能基于原始进行裁剪 (by `[]`) 拼接 (by `+`)，产生一个**新字符串**。
- 拼接时，会逐个**拷贝**到一个新的连续内存空间。如果不断 `+`，内存分配 pressure ↑。

:smile: ​`strings.Builder`

- 内置 []byte 对应一个底层数组。
- 使用 `unsafe.Pointer` 可以对内存直接进行操作，但 Builder **只允许拼接或重置**。
- 支持**自动扩容** x2。只要没有扩容，那么就不会发生拷贝。
- 支持**手动扩容**。
- **:warning: 一旦开始使用，就不可变。即不能被拷贝，也不能手动扩容。**

```go
var b strings.Builder
b.Grow(64) // expand 64 bytes more 

b.WriteString("Hello")
b.WriteString(" ")
b.WriteString("World")

b.String() // get content
b.Len()    // length

b.Reset()  // clear
```

:smile: `strings.Reader` 高效读取，内置计数器保存已读取得字节数。

`Len()` **表示的是未被读取字节的长度，而不是已存字节的长度。**

```go
reader := strings.Reader

buf := make([]byte, 4)
n, err := reader.Read(buf)
if err != nil && err != io.EOF {
	fmt.Println("Error reading:", err)
	return
}
fmt.Printf("Read %d bytes: %s\n", n, string(buf[:n]))


readingIndex := reader.Size() - int64(reader.Len()) // count of read
fmt.Printf("Reading index: %d", readingIndex)
```

`strings.Reader.ReadAt` 主要用于从数据流的某个**特定位置**开始读取数据，不改变流得当前位置。

```go
reader := strings.Reader
buf := make([]byte, 5)
n, err := reader.ReadAt(buf, 7) // starting from offset 7
if err != nil {
		fmt.Println("Error reading:", err)
		return
}
```

`strings.Reader.Seek` 主要用于在数据流中定位到一个新的位置，然后可以从该位置继续读取数据。**适合随机/反复读取**。

```go
reader := strings.Reader
_, err := reader.Seek(7, 0)
buf := make([]byte, 5)
n, err := reader.Read(buf)
```

**一次性读取**直到 `io.EOF`

```go
reader := strings.Reader
var result []byte
buf := make([]byte, 4)

// read 4 bytes per loop until EOF
for {
	n, err := reader.Read(buf)
	if err == io.EOF {
		break
	}
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	result = append(result, buf[:n]...)
}
```

### Common

| func                                   | Desc                             |
| -------------------------------------- | -------------------------------- |
| `len(str)`                             | 返回字符串长度 O(1)              |
| `range []rune(str)`                    | 遍历字符串转换成 rune/int32      |
| `strconv.Atoi(str)`                    | 数字字符串转整数                 |
| `strconv.Itoa(...)`                    | 整数转字符串                     |
| `[]byte(str)`                          | 字符串转字节                     |
| `string([]byte{...})`                  | 字节转字符串                     |
| `strconv.FormatInt(i int64, base int)` | 十进制数转2/8/16进制字符串       |
| `strings.Contains`                     | 判断是否包含子串                 |
| `strings.Count`                        | 统计字串出现次数                 |
| `str1 == str2`                         | 判断字符串是否相等               |
| `strings.EqualFold`                    | 判断字符串是否相等，不区分大小写 |
| `strings.Replace`                      | 替换字符串                       |
| `strings.ToUpper/ToLower`              | 大小写转换                       |
| `strings.Split`                        | 按 sep 进行拆分，返回一个切片    |
| `strings.Join`                         | 以 sep 进行组合切片成一个字符串  |
| `strings.Trim*`                        | 裁剪                             |
| `HasPrefix`<br />`HasSuffix`           | 是否有前后缀                     |