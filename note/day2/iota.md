iota：**特殊常量，可以认为是一个可以被编译器修改的常量**；表示 const 块中每个常量所处位置在块中的自动递增偏移值 offset. 每一行其值 ++。

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