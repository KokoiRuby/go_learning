`new`: 分配堆内存，给定一个**类型**，返回一个**指针/地址** → 该类型的 zero-value。

```go
func new(Type) *Type

var num *int
num = new(int)
fmt.Printf("%T\n", num)  // *int
fmt.Printf("%v\n", num)  // 0xc000014088
fmt.Printf("%v\n", &num) // 0xc00000e018
fmt.Printf("%v\n", *num) // 0
```

`make`: 分配堆内存，用于**初始化**并返回一个引用类型的**实例**：slice, map, channel

```go
func make(T, args)

// slice
s := make([]int, 5)
fmt.Println(s)

// map
m := make(map[string]string)
m["key"] = "value"
fmt.Println(m)

// channel
c := make(chan int, 1)
c <- 1
fmt.Println(<-c)
```

