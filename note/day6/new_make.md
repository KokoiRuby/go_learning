`new`: 分配内存(堆 Heap)，给定一个**类型**，返回一个**指针/地址** → 该类型的 zero-value,不会初始化这块内存

```go
func new(Type) *Type

var num *int
num = new(int)
fmt.Printf("%T\n", num)
fmt.Printf("%v\n", num)
fmt.Printf("%v\n", &num)
fmt.Printf("%v\n", *num)
```

`make`: 分配内存(堆 Heap)，用于**初始化**并返回一个引用类型的实例：slice, map, channel

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

