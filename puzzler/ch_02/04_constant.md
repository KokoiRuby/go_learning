### Constant

常量：**编译期被初始化**且可求解，**运行时其值不会被修改** ← ∵ 直接嵌入可执行文件中，**不会保存在内存里**。

类型仅限：bool, int, float, complex, string，即基本数据类型 & string。

```go
// type can be ignored as it can be inferred.
const identifier [type] = value

const s string = "abc"
const s = "abc"

// multi
const a, b, c = 1, false, "str"
```

```go
// enum
const (
    Unknown = 0
    Female = 1
    Male = 2
)
```

### iota

**特殊常量，可以认为是一个可以被编译器修改的常量**（编译器会自动帮忙计算）。

表示 const 块中每个常量所处位置在块中的**自动递增偏移值** offset → 每一行其值 ++。

每遇到一个新的常量块或单个常量声明时，都会重置为 0。

```go
// enum
const (
    a = iota // 0
    b = iota // 1
    c = iota // 2
)

// simple
const (
    a = iota // 0
    b        // 1 
    c        // 2
)
```



