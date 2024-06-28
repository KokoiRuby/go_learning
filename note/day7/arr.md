数组是具有**相同唯一类型**的一组带**索引/下表定长**的序列。通过索引/下标获取元素。

```go
var arr [size]type
```

```go
var balance [5]float32 // 0.0, 0.0, 0.0, 0.0, 0.0

var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

// assign to specific index
balance := [5]float32{1:2.0,3:7.0} // 0.0, 2.0, 0.0, 7.0, 0.0
```

数组**内存空间连续**。

```go
arr := [3]int{1, 2, 3}
// addr is consecutive
for i := 0; i < len(arr); i++ {
	fmt.Printf("%d - address is: %v\n", arr[i], &arr[i])
}

// v is a local var, its addr is fixed in each iter of loop
for _, v := range arr {
		fmt.Printf("%d - address is: %v\n", v, &v)
}
```

数组是**值类型**，∴ 赋值传参的时候是值拷贝 (浅)。如果要修改数组的值，需要传递指针/引用，注意参数列表中的长度要保持一致。

```go
func modify(arr *[3]int) {
	(*arr)[0] = 4 // * → value to modify
}

modify(&arr) // pass pointer
```

