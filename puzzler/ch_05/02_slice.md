切片是数组的一个**引用/视图**，所以是**引用类型**，传递效率更高。修改数组的值，会直接反映到切片上。

每个切片必定有一个底层数组，底层数组不可视。底层数组只有在没有任何切片指向的时候才会被 GC。

切片长度可变。

**通过引用数组创建**

```go
// ref. to arr
arr := [5]int{1, 2, 3, 4, 5}
sl := arr[1:3]
sl := arr[:]
sl := arr[2:] 
sl := arr[:4] 
```

**通过 `make` 进行创建**，其数组底层维护，不可操作。

```go
// []T, len, cap
var sl []int = make([]int, 5, 10)
```

**直接初始化创建**

```go
// len = cap = 3
var sl []int = []int {1,3,5}
```

**通过切片创建**

```go
arr := [5]int{1, 2, 3, 4, 5}
sl1 := arr[:]
sl2 := sl1[2:]
```

**内存结构**：

- ptr 指向底层数组的地址

- 切片长度/容量 len <= cap

  - `len(sl)` 获取长度
  - `cap(sl)` 获取容量


![Slice internal example[fig:Slice-internal-example]](https://www.practical-go-lessons.com/img/slice_inernal_example.6aa32199.png)

**遍历**：for | for range

```go
for i := 0; i < len(sl); i++ {}
for i, v := range sl {} 
```

**空切片**：未初始化之前为 nil，长度为 0

```go
var nilSlice []int
```

**传递给函数**：如果你有一个函数需要对数组做操作，可以传递数组的切片。

```go
var arr = [5]int{0, 1, 2, 3, 4}
myFun(arr[:])
```

**追加**：将**若干**具有相同类型的元素追加到切片后面并且**返回新的切片**。

- `append(slice []Type, elems ...Type)`
- 如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。
- 在底层创建一个新的数组 newArr 作扩容。
- 策略：
  1. len < 1024，cap *= 2
  2. len >= 1024，cap *= 1.25

**拷贝**：

- `copy(dst, src []Type)`
- 拷贝切片指向的是一个新的底层数组，所有修改不会影响原数组和切片。
- 多余的空间，零值填充。不够的空间则截断。

**UC**

- 将切片 b 的元素追加到切片 a 之后：`a = append(a, b...)`
- 复制切片 a 的元素到新的切片 b 上：`b = make([]T, len(a))` + `copy(b, a)`
- 删除位于索引 i 的元素：`a = append(a[:i], a[i+1:]...)`
- 切除切片 a 中从索引 i 至 j 位置的元素：`a = append(a[:i], a[j:]...)`
- 为切片 a 扩展 j 个元素长度：`a = append(a, make([]T, j)...)`
- 在索引 i 的位置插入元素 x：`a = append(a[:i], append([]T{x}, a[i:]...)...)`
- 在索引 i 的位置插入长度为 j 的新切片：`a = append(a[:i], append(make([]T, j), a[i:]...)...)`
- 在索引 i 的位置插入切片 b 的所有元素：`a = append(a[:i], append(b, a[i:]...)...)`
- 取出位于切片 a 最末尾的元素 x：x, `a = a[len(a)-1:], a[:len(a)-1]`
- 将元素 x 追加到切片 a：`a = append(a, x)`

