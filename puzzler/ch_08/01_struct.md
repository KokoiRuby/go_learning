Go 不支持纯粹的 OOP；没有 class，但支持 OOP 的特性，即可以通过一些方式实现类继承/封装/多态。

:bookmark_tabs: Review

- 封装（Encapsulation）：将**属性和方法**封装在一个类中，**限制**对内部数据和方法的**访问**，隐藏了实现的细节。
- 继承（Inheritance）：**子类**从**父类**继承属性和方法，子类也可以通过添加自己的特性和**重写**父类方法。
- 多态（Polymorphism）：指**同名方法**可以在**不同的对象上具有不同的行为**，灵活性和可扩展性 ↑。

使用 `new` 函数给一个新的结构体变量分配内存，它**返回**指向已分配内存的**指针**。

![img](https://cdn.learnku.com/uploads/images/201808/27/23/4BIeFlKTBD.jpg?imageView2/2/w/1240/h/0)

`p := Point{10, 20}` 表示创建一个类型为 `Point` 的结构体实例 `p`，并将字段 `x` 和 `y` 的值分别设置为 `10` 和 `20`。这种方式**直接创建了结构体的值类型实例**。

`p := &Point{10, 20}` 表示创建一个类型为 `*Point` 的结构体指针 `p`，并将指针指向一个新创建的 `Point` 结构体实例。这种方式**创建了结构体的指针类型实例**。

![img](https://cdn.learnku.com/uploads/images/201808/27/23/OLAUFPV0cu.jpg?imageView2/2/w/1240/h/0)

结构体名**首字母大写表示公有**，可被其他包访问。小写私有，只能在包内访问。

```go
type Cat struct {
	Name  string
	Age   int
	color string
	ptr   *int
	sl    []int
	m     map[string]string
}

// 1
var cat Cat
// 2
cat := Cat{}
// 3
var cat *Cat = new(Cat)
// 4
cat := new(Cat)
// 5
var cat *Cat = &Cat{...}
```

```go
type PublicStruct struct {
	Field1 string
	Field2 int
}

type privateStruct struct {
	Field1 string
	Field2 int
}

otherpackage.PublicFunction(p1) // ok
otherpackage.privateStruct(p2)  // nok
```

通过 `.` 选择器 selector 操作符来**访问成员**

```go
cat.Name = "Kitty"
fmt.Println(cat.Name)
```

结构体是**值类型**，传递值，修改不影响本体。

结构体若成员包含指针/slice/map，由于引用类型，其零值都是 nil，所以使用时要先 make。

```go
num := 1
cat.ptr = &num

cat.sl = make([]int, 3)
cat.sl = []int{1, 2, 3}

cat.m = make(map[string]string)
cat.m["key"] = "value"
```

结构体的所有字段在内存中是**连续分布**的，可以通过地址加减快速找到下一个。

如果字段是指针类型，那么他们也是连续的，但是指向的地址不一定是连续的。

```go
type Point1 struct {
	x, y int
}

type Rectangle1 struct {
	a, b Point1
}

r1 := Rectangle1{Point1{1, 2}, Point1{3, 4}}
fmt.Println(&r1.a.x, &r1.a.y)
fmt.Println(&r1.b.x, &r1.b.y)
```

```go
type Point2 struct {
	x, y int
}

type Rectangle2 struct {
	a, b *Point2
}
r2 := Rectangle2{&Point2{5, 6}, &Point2{7, 8}}
fmt.Println(&r2.a, &r2.b)
fmt.Println(r2.a, r2.b)
fmt.Println(*r2.a, *r2.b)
```

**结构体转换**：前提是两个结构体类型在**字段名、字段类型和字段顺序**上是**完全相同**的。

注意：如果给 type 起别名，那么 Go 也认为是两个不同的数据类型，需要显式转换。

```go
type Person struct {
	Name   string
	Age    int
	Salary float64
}

type Employee struct {
	Name   string
	Age    int
	Salary float64
}

p := Person{
	Name: "john",
	Age:  20,
}
// Person → Employee
var e Employee
e = Employee(p)
fmt.Println(e, p)
```

结构体 tag：meta-like in K/V 可通过**反射**获取，常用于**反序列化**。

```go
type Dog struct {
	Name  string `json:"Name"`
	Age   int    `json:"Age"`
	Color string `json:"Color"`
}

dog := Dog{"Bruce", 5, "White"}
nameTag := reflect.TypeOf(dog).Field(0).Tag.Get("json")
ageTag := reflect.TypeOf(dog).Field(0).Tag.Get("json")
colorTag := reflect.TypeOf(dog).Field(0).Tag.Get("json")
```

**匿名字段**：只指定类型，未指定字段名；**直接通过字段类型访问匿名字段的成员**，而无需通过字段名称进行层层嵌套。

```go
type D struct {
	string
	int
}

d := D{"Alice", 25}
fmt.Println(d.string)
fmt.Println(d.int)
```

可实现嵌入（类继承）。通过它，结构体可以**继承和访问其他结构体类型的字段和方法**。

:warning: 无论方法的签名是否一致，**外层**结构体中的方法都会**覆盖**被**嵌入**结构体中的同名方法。

```go
type Person struct {
	Name string
	Age  int
}

func (p *Person) Greet() {
	fmt.Println("Hello, I'm a person.")
}

type Employee struct {
	Person // embed
	Company string
}

// will overwrite Person.Greet()
func (e *Employee) Greet() {
	fmt.Println("Hello, I'm an employee.")
}

emp := Employee{
	Person: Person{
		Name: "John",
		Age:  30,
	},
	Company: "ABC Corp",
}

emp.Name    // John
emp.Age     // 30
emp.Company // ABC Corp
emp.Greet() // Hello, I'm an employee.
```

**匿名结构体**：在声明结构体变量时，直接定义结构体类型而不指定结构体名称。方便创建临时、只在某个作用域内使用的结构体对象。

```go
p := struct {
		Name  string
		Age   int
	}{
		Name:  "Alice",
		Age:   25,
	}
```

**结构体嵌入类型指针**：注意初始化的方式。

```go
type Person struct {
	Name string
}

func (p *Person) Greet() {
	fmt.Println("Hello, I'm a person.")
}

type Employee struct {
	*Person
	Company string
}

emp := &Employee{
		Person:  &Person{Name: "John"},
		Company: "ABC Corp",
}

emp.Greet()  // Hello, I'm a person.
emp.Name     // John
emp.Company  // ABC Corp
```

**空结构体**：`struct{}`，没有任何字段，不占任何空间。常用于管道元素类型。

```go
ch := make(chan struct{})
ch := make(chan struct{}, bufferSize)

go func() {
    ch <- struct{}{} // send to
}()

// recv from
<-ch
```

**结构体实例内存大小**：`size := unsafe.Sizeof(instnce)`

**结构体打印**：`fmt.Printf("%+v", struct)`

**结构体字段自引用**：只能用指针

```go
type Node struct {
    val int
    nxt *Node
}
```

