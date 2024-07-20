为了保护一个对象的并发修改，我们可以**使用一个后台的协程来顺序执行一个匿名函数**，而不是通过同步锁。

结构体中包含了匿名函数类型的通道。

```go
type Person struct {
	Name   string
	salary float64
	chF    chan func()
}
```

构造器中用一个协程启动一个 `backend()` 方法。

该方法在一个无限 for 循环中执行所有被放到 chF 上的函数，有效的序列化他们，从而提供安全的并发访问。

```go
func NewPerson(name string, salary float64) *Person {
	p := &Person{name, salary, make(chan func())}
	go p.backend()
	return p
}

func (p *Person) backend() {
	for f := range p.chF {
		f()
	}
}
```

setter/getter 执行时，会将匿名函数传递。无限循环的 backend() 就会执行。

**而通道保证了顺序执行，从而确保了并发安全。**

```go
func (p *Person) SetSalary(sal float64) {
	p.chF <- func() { p.salary = sal }
}

func (p *Person) GetSalary() float64 {
	fChan := make(chan float64)
	p.chF <- func() { fChan <- p.salary }
	return <-fChan
}
```

