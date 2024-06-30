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

