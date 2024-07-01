Go 接口也是一种**数据类型**。不需要显式实现，只需要一个变量**实现接口类型中的所有方法**，那么该变量就实现了这个接口。接口是一种**规范**，使用接口必须要按照它的规范来。

:confused: **vs?**

**继承 is-a，接口 like-a**

继承提高复用性/扩展性；接口在于规范的定义。

接口比继承更加灵活，且实现代码解耦。

声明

- 接口中所有的方法都没有方法体。
- 接口本身不能创建实例，但可以指向实现其接口的自定义类型变量。

```go
type interface_name interface {
   method_name1 [return_type]
   ...
}
```

多接口

```go
type A interface{
    Hello()
}

type B interface{
    Say()
}

type S struct {

}

func (s S) Helo() {}
func (s S) Say() {}

var s S
var a A = s
var b B = s
a.Hello()
b.Say()
```

空接口 nil：任何类型都可以实现空接口，或者说空接口可以代表任何类型。

```go
var a interface{}
a = 20
a = "hello"
a = true

var m = make(map[string]interface{})

var slice1 = []interface{}{1, 2, 3, 4, 5, "hello", true}
```

接口继承/嵌套，方法不能重名。

```go
type A interface {
    test01()
    // test04()  // nok
}

type B interface {
    test02()
    // test04()  // nok
}

type C interface {
    A
    B
    test03()
}
```

接口是引用类型，传递引用。

```go
type Usb interface {
    Say()
}

type Stu struct {
}

func (this *Stu) Say() {
    fmt.Println("Say something...")
}

stu := Stu{}
// nok: Stu does not implement Usb (method Say has pointer receiver)
// var u Usb = stu 
var u Usb = &stu
```
