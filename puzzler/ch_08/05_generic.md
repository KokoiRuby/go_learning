泛型 Generic since v1.18+，通过类型参数 `T` 限制具体的类型。

泛型接口

```go
type List[T any] interface {
	Add(index int, val T)
	Append(val T)
	Delete(index T)
}
```

泛型结构体

```go
type Node1[T any] struct {
	val T
}
```

泛型函数

```go
func Sum[T Number](vals ...T) T {
	var sum T
	for _, val := range vals {
		sum = sum + val
	}
	return sum
}

// ++constraint
type Number interface {
	int | uint
}
```

