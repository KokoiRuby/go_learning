Go 接口可以实现多态 Polymorphism。

多态允许**不同的类型以一致的方式进行处理**，即使这些类型实现了**不同的行为**。

通常我们可以将实现接口的类型统一放入一个接口类型的数组中。

```go
type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	name string
}

func (p Phone) Start() {
	fmt.Println("Starting...")
}

func (p Phone) Stop() {
	fmt.Println("Stopping...")
}

type Camera struct {
	name string
}

func (c Camera) Start() {
	fmt.Println("Starting...")
}

func (c Camera) Stop() {
	fmt.Println("Stopping...")
}

var usbArr [2]Usb
usbArr[0] = Phone{"iPhone"}
usbArr[1] = Camera{"Nikon"}

for _, elem := range usbArr {
	elem.Start()
	elem.Stop()
}
```

**指针接收者实现接口**：接口变量保存的是一个指针，这个指针指向了实现该接口的自定义类型的变量。

```go
type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

func (c *cat) move() {
	fmt.Println("Moving...")
}

func (c *cat) eat(food string) {
	fmt.Println("Eating", food)
}

var cat = &cat{"Tom", 4}
// var cat = cat{"Tom", 4} // nok
var animal = cat
animal.move()
animal.eat("Fish")
```

