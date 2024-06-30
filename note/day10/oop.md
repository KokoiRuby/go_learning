OOP 步骤：

1. 声明结构体
2. 编写结构体字段
3. 编写结构体方法

```go
type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}
```

创建结构体变量时，可直接指定字段值，不依赖定义顺序。

```go
stu3 := Student{
	age:  18,
	name: "Mary",
}
```

如果不指定字段名，必须要按照顺序，且必须填写所有的参数

```go
// nok: Too few values
stu4 := Student{"Bruce", 18}
```

