## Pass by *

> 描述变量在函数传递时的行为

**Pass by Value 传值**：函数接收的是该变量的**副本**，即堆内存中的变量值复制一份到函数调用栈内存中。对于基本数据类型 (integer/float/complex/bool/string/rune/byte/array) 变量，传递时都是传值，函数内对该变量的修改**不会影响**到函数外的原始变量。

```go
func modifyValue(val int) {
    val = 10
}

a := 5
modifyValue(a)
fmt.Println(a)  // 5
```

**Pass by Ref. 传引用**：函数接收的是该变量的**地址/指针**，即**内存地址**。对于引用类型 (slice/map/chan) 变量传递时传递的是引用，函数内对该引用指向的内容修改，**会影响**原始变量。 

```go
func modifyReference(val *int) {
    *val = 10
}

a := 5
modifyReference(&a)
fmt.Println(a)  // 10
```

## Receiver

> **函数调用**是全局作用域的，而**方法调用**是**绑定**到特定**类型**（如结构体）的**实例**上的。

**值类型**：方法会对接收者 (实例) 的一个**副本**进行操作，方法内部的修改**不会影响**到原始对象的状态。

通常用于**数据量较小或方法不需要修改接收者**的情况。

```go
type Rectangle struct {
    width, height int
}

func (r Rectangle) Area() int {
    return r.width * r.height
}

// 值类型接收者修改接收者的属性
func (r Rectangle) DoubleWidth() {
    r.width *= 2
}

rect := Rectangle{10, 5}
fmt.Println("Area:", rect.Area()) // Area: 50

rect.DoubleWidth()
fmt.Println("Width after DoubleWidth:", rect.width) // Width after DoubleWidth: 10
```

**引用类型**：方法接收者指向的是对象的**内存地址**，方法对接收者的操作**会影响**到原始对象的状态。

适用于需要在方法中**修改接收者状态的情况，或数据量较大，拷贝成本较高时**使用。

```go
type Rectangle struct {
    width, height int
}

func (r *Rectangle) Area() int {
    return r.width * r.height
}

func (r *Rectangle) DoubleWidth() {
    r.width *= 2
}

rect := Rectangle{10, 5}
fmt.Println("Area:", rect.Area()) // Area: 50

rect.DoubleWidth()
fmt.Println("Width after DoubleWidth:", rect.width) // Width after DoubleWidth: 20
```

## Copy

**浅复制 Shallow Copy**: 只复制对象本身，而不复制它所引用的其他对象。浅复制创建的新对象与原始对象**共享相同的子对象引用**，原始对象的某个子对象发生变化，复制新对象也会体现出来。

:smile: 快速内存高效 ∵ 只涉及指针复制。

适用于**对象内部没有引用类型数据或不需要对子对象进行修改**的场景。

```go
type Person struct {
    Name    string
    Friends []string
}

original := Person{
    Name:    "Alice",
    Friends: []string{"Bob", "Charlie"},
}

// shallow
copy := original

// modify original
original.Friends[0] = "David"

fmt.Println("Original:", original.Friends) // [David Charlie]
fmt.Println("Copy:", copy.Friends)         // [David Charlie]
```

**深复制 Deep**: **深复制**不仅复制对象本身，还**递归复制**它所引用的所有对象。深复制会创建一个全新的对象，并且这个新对象中所有的子对象也是新创建的，与原始对象完全独立。修改**不会影响**到彼此。

:cry: 复制会消耗更多的内存和时间，∵ 它需要递归地复制所有嵌套的对象。

适用于**对象内部包含引用类型数据，并且需要在复制后保持对象之间完全独立**的场景。

```go
type Person struct {
    Name    string
    Friends []string
}

// impl manually
func deepCopyPerson(p Person) Person {
    // new sl & copy content
    friendsCopy := make([]string, len(p.Friends))
    copy(friendsCopy, p.Friends)

    // a new obj
    return Person{
        Name:    p.Name,
        Friends: friendsCopy,
    }
}

original := Person{
    Name:    "Alice",
    Friends: []string{"Bob", "Charlie"},
}

copy := deepCopyPerson(original)

original.Friends[0] = "David"

fmt.Println("Original:", original.Friends) // [David Charlie]
fmt.Println("Copy:", copy.Friends)         // [Bob   Charlie]
```

