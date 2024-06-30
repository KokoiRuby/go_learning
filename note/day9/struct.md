Go 不支持纯粹的 OOP；没有 class，但支持 OOP 的特性，即可以通过一些方式实现类继承/封装/多态。

:bookmark_tabs: Review

- 封装（Encapsulation）：将属性和方法封装在一个类中，限制对内部数据和方法的访问，隐藏了实现的细节。
- 继承（Inheritance）：子类从父类继承属性和方法，子类也可以通过添加自己的特性和重写父类方法。
- 多态（Polymorphism）：指同名方法可以在不同的对象上具有不同的行为，灵活性和可扩展性 ↑。



结构体**成员首字母大写表示公有**，可被其他包访问。小写包私有，只能在包内访问。

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

通过 `.` 操作符来**访问成员**

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

结构体的所有字段在内存中是连续分布的，可以通过地址加减快速找到下一个数字。

如果时字段是指针类型，那么他们也是连续的，但是指向的地址不一定是连续的。

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

结构体转换：前提是两个结构体类型在字段名、字段类型和字段顺序上是**完全相同**的。

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

结构体 tag：meta-like in K/V 可通过**反射**获取，常用于反序列化。

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

