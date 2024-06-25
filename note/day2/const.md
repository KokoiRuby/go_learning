常量：编译期就被初始化且可求解，运行时其值不会被修改。

类型仅限：bool, int, float, complex, string.

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

