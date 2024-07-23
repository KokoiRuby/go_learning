Go 接口也是一种**数据类型**。不需要显式实现，只需要一个变量**实现接口类型中的所有方法**，那么该变量就实现了这个接口。

接口是一种**规范**，使用接口必须要按照它的规范来。

接口无法被实例化 neither `new` nor `make`

**vs**

- **继承 is-a，接口 like-a**
- 继承提高复用性/扩展性；接口在于规范的定义。
- 接口比继承更加灵活，且实现代码解耦。

**声明**

- 接口中所有的方法都没有方法体。
- 接口本身不能创建实例，但可以指向实现其接口的自定义类型变量。
  - 只有自定义数据类型的**所有**方法名和签名同接口定义的一直，就说明该类型实现了对应的接口。

```go
type interface_name interface {
   method_name1 [return_type]
   ...
}
```

**多接口**

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

**空接口 nil**：任何类型都可以实现空接口，或者说空接口可以代表任何类型。

:warning: 将某个类型切片赋值到一个空接口切片，只能通过 for-range 显式赋值。

```go
var a interface{}
a = 20
a = "hello"
a = true

var m = make(map[string]interface{})

var slice1 = []interface{}{1, 2, 3, 4, 5, "hello", true}
```

```go
var dataSlice []myType = FuncReturnSlice()
var interfaceSlice []interface{} = make([]interface{}, len(dataSlice))
for i, d := range dataSlice {
    interfaceSlice[i] = d
}
```

**接口继承/嵌套**：方法不能重名。

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

**接口组合**：接口类型作匿名成员列表即可。

:construction_worker: 每个接口尽量小，然后相互组合。

```go
type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string)
}

// combo
type ReadWriter interface {
	Reader
	Writer
}
```

接口是**引用类型**，传递引用。

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

**赋值接口变量**：传递的是**副本**；其中接口变量的类型为**静态类型**，而赋值的类型为**动态类型**。

动态类型信息保存在 iface 中包含 2 个指针：指向动态类型 & 指向动态值。

```go
type Pet interface {
    Name() string
    Category() string
}

type Dog struct {
	name string
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) Category() string {
	return "dog"
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

dog := Dog{"little pig"}
var pet Pet = dog        // pet's static type: Pet, dynamic type: Dog
dog.SetName("monster")   // pet.name won't get changed
```

**接口变量值为 nil**：只有在声明时不进行初始化 or 将 `nil` 赋值。

```go
var pet Pet        // nil
var pet Pet = nil  // nil

var dog Dog
var pet Pet = dog  // not nil
```

**把一个值为 nil 的某个实现类型的变量赋给了接口变量，那么在这个接口变量上仍然可以调用该接口的方法吗？**

```go
var dog Dog
var pet Pet = dog
pet.SetName("Puppy") // panic: runtime error: invalid memory address or nil pointer dereference
```

**测试某个值是否实现了某个接口？**

```go
type myInt interface {
   ...
}

if sv, ok := v.(myInt); ok {
    // TODO
}
```

### Why

`interface `类型存在的意义是为了实现 OOP 思想 → “高内聚/低耦合”。

**“开闭原则”：扩展开放，修改关闭。**

:cry: 增加方法满足业务，导致模块越来越臃肿

- Banker: [save() + pay() + transfer()] + share(), membership()...

:smile: 抽象出模块，不同的功能对应不同的实现 = **多态思想：通过父类指针 → 子类实例**

- AbstractBanker: DoBusi()

  - SaveBanker: DoBusi()

  - PayBanker: DoBusi()

  - TransferBanker: DoBusi()

    ++

  - ShareBanker: DoBusi()

  - MembershipBanker(): DoBusi()

**“依赖倒转”：面向接口/抽象编程。**

:cry: [A|B].DriveBenz, [A|B].DriveBMW, [A|B]DriveTesla ← ++C & Toyota

:smile: [A|B|C] → <u>Driver.DriveCar</u> ← Benz, BMW, Tesla, Toyota
