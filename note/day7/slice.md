切片是数组的一个**引用/视图**，所以是引用类型。修改数组的值，会直接反映到切片上。

每个切片必定有一个底层数组。

切片长度可变。



通过引用数组创建

```go
// ref. to arr
arr := [5]int{1, 2, 3, 4, 5}
sl := arr[1:3]
sl := arr[:]
sl := arr[2:] 
sl := arr[:4] 
```

通过 `make` 进行创建，其数组底层维护，不可操作。

```go
var sl []int = make([]int, 5, 10)
```

直接初始化创建

```go
// len = cap = 3
var sl []int = []int {1,3,5}
```

还可以通过切片创建

```go
arr := [5]int{1, 2, 3, 4, 5}
sl1 := arr[:]
sl2 := sl1[2:]
```



内存结构：

- ptr 指向数组的地址

- 切片长度/容量 len <= cap

  - `len(sl)` 获取长度
  - `cap(sl)` 获取容量

  

![Slice internal example[fig:Slice-internal-example]](https://www.practical-go-lessons.com/img/slice_inernal_example.6aa32199.png)



遍历：for | for range

```go
for i := 0; i < len(sl); i++ {}
for i, v := range sl {} 
```



空切片：未初始化之前为 nil，长度为 0

```go
var nilSlice []int
```



追加：

- `append(slice []Type, elems ...Type)`
- 如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。
- 在底层创建一个新的数组 newArr 作扩容，同样底层维护，不可视。
- 策略：
  1. len < 1024，cap *= 2
  2. len >= 1024，cap *= 1.25



拷贝：

- `copy(dst, src []Type)`
- 拷贝切片指向的是一个新的底层数组，所有修改不会影响原数组和切片。
- 多余的空间，零值填充。不够的空间则截断。