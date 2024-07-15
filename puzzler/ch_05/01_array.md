数组是具有**相同唯一类型**的一组带**索引/下表定长**的序列。通过索引/下标获取元素。

可以使用**空接口类型**，但使用值的时候要进行类型断言。

数组元素可变，支持通过下标修改值。

```go
var arr [size]type
```

```go
var balance [5]float32 // 0.0, 0.0, 0.0, 0.0, 0.0

var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

// assign to specific index
balance := [5]float32{1:2.0,3:7.0} // 0.0, 2.0, 0.0, 7.0, 0.0

// set val by idx
balance[0] = 2.3

// new
arr := new([5]int)
```

数组**内存空间连续**。范围：`0 ~ len(arr)-1`

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

数组是**值类型**，∴ 赋值传参的时候是值拷贝(浅)。如果要在函数内修改数组的值，需要传递指针/引用，注意参数列表中的**长度要保持一致。**

:construction_worker: 对于大数组，传递会消耗很多内存；可以传递数组的指针 or 数组的切片。（一般不这么做，直接使用切片即可）

```go
func modify(arr *[3]int) {
	(*arr)[0] = 4 // * → value to modify
}

modify(&arr) // pass pointer
```

多维数组 matrix

声明

```go
var matrix [SIZE1][SIZE2]...[SIZEN] type
```

```go
matrix := [2][3]int{
	{1, 2, 3},
	{4, 5, 6},
}

numbers := [][]int{}
row1 := []int{1, 2, 3}
row2 := []int{4, 5, 6}
row3 := []int{7, 8, 9}
numbers = append(numbers, row1, row2, row3)
```

遍历

```go
for i, v1 := range matrix {
	for j, v2 := range v1 {
		fmt.Printf("matrix[%v][%v] = %v \n", i, j, v2)
	}
}
```