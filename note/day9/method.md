Go 方法 Method 是与特定类型关联的函数（不能独立存在的函数）。通过方法，允许在**自定义类型**上定义操作（访问字段或执行响应的逻辑）。

方法通过接收者（receiver）来与类型进行关联，**接收者可以是结构体或非结构体**。

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

调用和传参：不同的于函数的是，方法调用时会**将调用方法的变量作为接收者传递给方法**。

- 如果接收者是值类型，那么是传递值。（方法内修改不会影响原变量）
  - **我们也可以用指针类型的变量进行调用，但编译器帮我们简化了**
- 如果接收者是引用类型，那么传递引用。（方法内修改会影响原变量）

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

:confused: **vs**

声明不同：方法多一个接收者。

调用不同：方法通过类型的实例来调用。函数可以直接通过函数名调用。

声明位置：方法必须声明在同一个包内；函数可以声明在任何包中。

