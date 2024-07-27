### OOP

步骤：

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

### Abstract

抽象 Abstraction：把一类事物的**共有属性(字段)和方法**提取出来，**形成一个模型(模板)**。

```go
type Account struct {
	AccountNo string
	Password  string
	Balance   float64
}

func (account *Account) Deposite(money float64, pwd string) {}
func (account *Account) WithDraw(money float64, pwd string) {}
func (account *Account) Query(pwd string) {}
```

### Encapsulation

封装：将数据和操作封装在一个类中，通过限制对内部数据和方法的访问，隐藏了实现的细节。封装提供了数据的安全性和保护性，外部只能通过公共接口访问数据和方法。seter/getter-like

```go
// /encapsulation/model/person.go
package modedl

import {
    "fmt"
}

type person struct {
    Name string
    age int
    sal float64
}

// factory as constructor
func newPerson(name String) *person{
    return &person{
        Name: name,
    }
}

// seter/getter for pri fields
func (p *person) SetAge(age int) {
    p.age = aget
}

func (p *person) GetAge() int {
    return p.age
}

func (p *person) SetSal(sal float64) {
    p.sal = sal
}

func (p *person) GetSal() float64 {
    return p.sal
}
```

```go
// /encapsulation/main.go
package main
import (
	"fmt"
	"encapsulate/model"
)

func main() {
    p := model.NewPerson("Jack")
    p.SetAge(18)
    p.SetSal(5000)
    fmt.Println(p)
    fmt.Println(p.GetAge(), pGetSal())
}
```

### Inheritance

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

父类方法无论时公有还是私有，子类实例都可以直接调用。一般可简化，把父类去掉。

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

父子类字段重复时，就近原则。

```go
type A struct{
	Name string
	age int
}

type B struct{
    Name string
    A
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

### Composition

**鸭子类型**：态类型语言中的一种类型风格，对象的适用性不是基于继承特定的类或实现接口，而是取决于对象的方法和属性。

Go 以**组合**的方式实现鸭子类型。

- 接口组合接口：接口扩展
- 结构体组合结构体：类继承
- 结构体组合结构体指针：类继承
- 结构体组合接口：结构体实现方法

:construction_worker: A 组合 B 后，可直接在 A 上调用 B 方法；B 已实现的接口，A 也相应实现；

```go
type Inner struct {
}

func (i Inner) DoSomething() {
    fmt.Println("Inner")
}

type Outer struct {
	Inner
}

var o Outer
// same
o.DoSomething()        // Inner
o.Inner.DoSomething()  // Inner
```

:construction_worker: 同名方法 B 会覆盖 A

```go
type Inner struct {
}

func (i Inner) DoSomething() {
    fmt.Println("Inner")
}

type Outer struct {
	Inner
}

func (o Outer) DoSomething() {
    fmt.Println("Outer")
}

o.DoSomething()        // Outer
o.Inner.DoSomething()  // Inner
```

:construction_worker: 注意！

```go
type Inner struct {
}

func (i Inner) SayHello() {
	fmt.Println("Hello, " + i.Name())
}

func (i Inner) Name() string {
	return "Inner"
}

type Outer struct {
	Inner
}

func (o Outer) Name() string {
	return "Outer"
}

o.SayHello() // Hello Inner
```

