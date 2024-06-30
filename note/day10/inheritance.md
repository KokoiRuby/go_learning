继承：子类从父类继承属性和方法，并可以添加自己的特性或重写父类的方法，实现代码重用和扩展。

Go 中通过 struct 嵌套 Embedding 实现（类）继承。

```go
type parent struct {
    field1 type
    field2 type
    ...
}

type child struct {
    parent
    field3 type
}
```

父类方法无论时公有还是私有子类实例都可以直接调用。一般可简化，把父类去掉。

```go
var b B
b.A.Name = "Jack"
b.Name = "Tom"
b.A.age = 18
b.age = 18

b.A.SayOk()
b.SayOk()
b.A.hello()
b.hello()
```

父子类字段重复时，就近原则

```go
type A struct{
	Name string
	age int
}

type B struct{
    A
    Name string
}

b.Name = "Jack"  // → B
b.A.Name = "Tom" // → A
```

如果嵌入多个结构体，且具有相同的字段和方法，必须指定是哪个父类，否则编译不过（毕竟不知道到底哪个）

```go
type A struct{
	Name string    
	age int
}
type B struct{
    Name string    
    score int
}
type C struct{
	A
	B
}

c.Name = "Jack"   // nok
c.A.Name = "Tom"
c.B.Name = "Mary"

c := C{
    A{"Tom", 18},
    B{"Mary", 1000},
}
```

