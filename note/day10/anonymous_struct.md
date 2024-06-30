**匿名结构体**：在声明结构体变量时，直接定义结构体类型而不指定结构体名称。

通过这种方式，可以方便地创建临时的、只在某个作用域内使用的结构体对象，而无需为其定义具体的结构体类型名称。

```go
p := struct {
		Name  string
		Age   int
		Email string
	}{
		Name:  "Alice",
		Age:   25,
		Email: "alice@example.com",
	}
```

**匿名字段**：声明结构体时，将字段的名称省略，只指定字段的类型。

可直接通过字段类型访问匿名字段的成员，而无需通过字段名称进行层层嵌套。

```go
type D struct {
	string
	int
}

d := D{"Alice", 25}
fmt.Println(d.string)
fmt.Println(d.int)
```

