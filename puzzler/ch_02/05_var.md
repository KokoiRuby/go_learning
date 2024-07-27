#### Variable

变量：可视为一个内存块，用于存储**特定类型** (明确边界) 的数据，并可以通过**变量名**进行访问和操作。

```go
var identifier type
// multiple
var identifier1, identifier2 type
```

声明

1. **指定变量类型，如果没有初始化，则变量默认为零值 Zero-value**

```go
var v int
fmt.Println(v) // 0

var v bool
fmt.Println(v) // false

var v string
fmt.Println(v) // ""

// nil
var v *int
var v []int
var v map[string] int
var v chan int
var v func(string) int
var v error
```

2. **根据值自行推导变量类型**

```go
// type is: int
var v = 13
```

3. **短变量声明 `:=`**

   :warning: 不可以再次对于相同名称的变量使用短变量声明，毕竟已经声明过一次了，后续只要赋值就行。
   
   :warning: 只能用于局部变量 & 方法内部，左边必须有一个新变量。

```go
v := 1
// equals to
var v int
v = 1

// nok
v := 2
// ok
v = 2
```

多变量声明

```go
v1, v2, v3 := 1, 2, 3 
// equals to
var v1, v2, v3 int
v1, v2, v3 = 1, 2, 3 

// global var
var (
    v1 int
    v2 string
)
```

#### vs

**Value Types** 值类型：int, float, bool, string

- 指向**堆内存 Heap ↑ new/malloc & delete/free** 中的值
- 当使用 `=` 将一个变量的值赋值给另一个变量时，实际上是将内存中的值进行**(浅)拷贝**，两个完全独立的内存区域，修改其中一个不会影响另一个。

**Reference Types** 引用类型：pointer, slice, map, channel, interface

- 指向值所保存的**地址** = 指针
- 当使用 `=` 将一个变量的值赋值给另一个变量时，**(深)拷贝**的是地址，两个地址指向同一个值。修改其中一个会直接影响另一个。



![Understanding Value Types and Reference Types in Java | by Enes Ekşi |  Medium](https://miro.medium.com/v2/resize:fit:706/1*-iKAezhqio6aPLMFU7oeXQ.jpeg)



<img src="https://i.sstatic.net/Ufj7o.png" alt="enter image description here" style="zoom: 67%;" />

