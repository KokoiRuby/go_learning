Go 方法 Method 是**与特定类型关联的函数**（不能独立存在的函数）。允许在**自定义类型**上定义操作（访问字段或执行响应的逻辑）。

- 方法必须命名。
- 不能被当作值进行传递。
- 必须属于某个类型，通过**接收者 receiver** 来与类型进行关联，**接收者可以是结构体或非结构体**。不能是接口和指针类型。
- 一个**数据类型关联的所有方法**，共同组成了该类型的方法集合。
- 同一个方法集合中的方法不能出现重名。但基于不同接收者类型是支持重载的。

```go
func (recv_name recv_type) method_name(param_list) (return_list) {
	// body
}
```

```go
type A struct{
    Age int
}

// struct A has a method named "test"
// this "test" binds to struct A
func (a A) test() {
    fmt.Println(a.Age)
    fmt.Println((&a).Age)   // compiler makes a → (&a)
}
```

```go
// overload on different receivers
func (r *Receiver_A) Add(n int) int
func (r *Receiver_B) Add(n int) int
```

如果方法不使用用 recv 的值，可以用 `_` 

```go
type MyStruct struct {
	value int
}

func (_ MyStruct) PrintSomething() {
	fmt.Println("Something")
}
```

**调用和传参**：不同的于函数的是，方法调用时会**将调用方法的变量作为接收者传递给方法**。

- 如果接收者是值类型，那么是传递值。（方法内修改不会影响原变量）
  - **我们也可以用指针类型的变量进行调用，但编译器帮我们简化了**
- 如果接收者是引用类型，那么传递引用。（方法内修改会影响原变量）

:construction_worker: 如果想要方法改变接收者的数据，就在接收者的引用类型上定义该方法。

:bookmark_tabs: **总结**

- 指针方法可以通过指针调用。
- 值方法可以通过值调用。
- 接收者是值的方法可以通过指针调用，因为指针会首先被解引用。
- 接收者是指针的方法不可以通过值调用，因为存储在接口中的值没有地址。

```go
// pass value
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// pass reference
func (r *Rectangle) Scale(factor float64) {
	r.width *= factor
	r.height *= factor
}
```

**func vs. method**

声明不同：方法多一个接收者。

调用不同：方法通过类型的**实例调用**。函数可以直接通过**函数名调用**。

声明位置：方法必须声明在同一个包内；函数可以声明在任何包中。

**getter/setter**：访问/修改结构体的私有字段。

```go
type Boss struct {
	name string
}

func (p *Boss) GetName() string {
	return p.name
}

func (p *Boss) SetName(name string) {
	p.name = name
}
```

`String()` 方法**定制化**字符串输出，会在 `fmt.Print[f|ln]` 中默认输出。对应 `%v`。

:warning: 不要在 `String()` 方法里面调用涉及 `String()` 方法的方法，否则无线递归。

```go
func (p *Person) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.name, p.age)
}
```

