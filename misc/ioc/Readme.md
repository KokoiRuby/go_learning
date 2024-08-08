## IoC

**Inversion of Control 反转控制**：一种软件设计原则，将控制权从代码内部转移给外部容器或框架，以实现松散耦合和可重用性。

**最常见的实现：DI Dependency Injection 依赖注入**，通过外部将依赖关系传递给组件，而不是组件自己创建或查找依赖关系。

思想：将控制逻辑 & 业务逻辑分离，**业务逻辑依赖控制逻辑** = **控制逻辑注入业务逻辑** = 将控制逻辑转换成接口，让业务逻辑去实现。

例子：

- 开关是控制逻辑，电器是业务逻辑，不要在电器中实现开关，而是把开关抽象成一种协议，让电器都依赖之。
- 如果某个类需要一个成员对象，那么这个对象不应该是类自己生成，而是由外部传入类构造器。

### Composite

Go 中通过**组合 Composite** 将一个结构体嵌入另一个。

```go
type Widget struct {
    X, Y int
}
type Label struct {
    Widget        // Embedding (delegation)
    Text   string // Aggregation
}
```

```go
label := Label{Widget{10, 10}, "State:"}
label.X = 11
label.Y = 12
```

```go
// Button ← Label ← Widget
type Button struct {
    Label // Embedding (delegation)
}

// ListBox ← Widget
type ListBox struct {
    Widget          // Embedding (delegation)
    Texts  []string // Aggregation
    Index  int      // Aggregation
}
```

通过接口 interface 实现方法重写

```go
type Painter interface {
    Paint()
}
 
type Clicker interface {
    Click()
}
```

实现：

- Label            → Painter
- Button & ListBox → Painter & Clicker

```go
func (label Label) Paint() {
  fmt.Printf("%p:Label.Paint(%q)\n", &label, label.Text)
}

// if not implemented, will call Label.Paint()
// if implemented, overriding
func (button Button) Paint() {
    fmt.Printf("Button.Paint(%s)\n", button.Text)
}
func (button Button) Click() {
    fmt.Printf("Button.Click(%s)\n", button.Text)
}


func (listBox ListBox) Paint() {
    fmt.Printf("ListBox.Paint(%q)\n", listBox.Texts)
}
func (listBox ListBox) Click() {
    fmt.Printf("ListBox.Click(%q)\n", listBox.Texts)
}
```

多态：父类指针/引用 → 子类的具体实现

```go
button := Button{Label{Widget{10, 70}, "OK"}}
listBox := ListBox{Widget{10, 40}, []string{"AL", "AK", "AZ", "AR"}, 0}

for _, painter := range []Painter{button, listBox} {
    painter.Paint()
}
```

### Practice

`IntSet` is buz while `Undo` is logic control.

```go
// bad
type IntSet struct {
	data map[int]struct{}
}

type UndoableIntSet struct {
	IntSet                 // buz in logic
	functions []func()
}


// good
type Undo []func()

func (undo *Undo) Add(function func()) { ... }
func (undo *Undo) Undo() error { ... }

type IntSetNew struct {
	data map[int]bool
	undo Undo // dep inject, logic in buz
}
```



