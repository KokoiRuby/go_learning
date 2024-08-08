## MapReduce

MapReduce 是一种编程模型和处理大规模数据集的计算框架/控制逻辑。

MapReduce 由 Google 提出，用于**并行化和分布式处理大规模数据集**，以便能够高效地在大量计算机集群上运行。

![img](https://media.geeksforgeeks.org/wp-content/uploads/20230420231217/map-reduce-mode.png)

流程：

1. **Map**：
   - 输入数据集被划分成若干个小的子集，每个子集称为一个 K/V。
   - **并行执行**，每个输入 K/V 经过 Map 函数处理后，生成若干个中间 K/V。
2. **Shuffle/Filter**：
   - 过渡阶段，中间 K/V 会被重新**分区和排序**，以便将**相同 K** 的数据传递给同一个 Reduce。
3. **Reduce**：
   - 接收中间 K/V，并对具有**相同 K 的 V** 进行**归并**操作。

UC：

- 大数据处理领域被广泛应用，例如数据分析、日志处理、搜索引擎索引构建

- Hadoop 是一个基于 MapReduce 的开源框架，提供了对大规模数据的存储和处理能力。

Example:

```go
USER_ID MOVIE_ID SCORE TIMESTAMP
196     242      3     881250949
186     302      3     891717742
196     377      1     878887116
244     51       2     880606923
166     346      1     886397596
186     474      4     884182806
186     265      2     881171488
```

1. Map:

```go
// k → v
196:242
186:302
196:377
244:51
166:346
186:27
186:265
```

2. Shuffule/Filter:

```go
// merget
166:346
186:302,274,265
196:242,377
244:51  
```

3. Reduce:

```go
// count
166:1
186:3
196:2
244:1 
```

### Map

```go
// MapStrToInt iter a string arr & call fn then append to a new int arr
func MapStrToInt(arr []string, fn func(s string) int) []int {
	var newArray []int
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}

list = []string{"1", "2", "3"}
y := MapStrToInt(list, func(s string) int {
	i, _ := strconv.Atoi(s)
	return i
})
fmt.Printf("%v\n", y) //[1, 2, 3]
```

### Shuffle/Filter

```go
// Filter iter int arr & call fn to filter then append to a new int arr
// fn returns bool as cond
func Filter(arr []int, fn func(n int) bool) []int {
	var newArray []int
	for _, it := range arr {
		if fn(it) {
			newArray = append(newArray, it)
		}
	}
	return newArray
}

var s = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
out := Filter(s, func(n int) bool {
	return n%2 == 1
})
fmt.Printf("%v\n", out) // [1 3 5 7 9]
```

### Reduce

```go
// Reduce iter string arr & call fn to then sum
func Reduce(arr []string, fn func(s string) int) int {
	sum := 0
	for _, it := range arr {
		sum += fn(it)
	}
	return sum
}

var list = []string{"a", "bb", "ccc"}
x := Reduce(list, func(s string) int {
	return len(s)
})
fmt.Printf("%v\n", x) // 6
```

### Essense

Map/Reduce/Filter只是一种**控制逻辑**，**真正的业务逻辑是在传给他们的数据和那个函数来定义的**，经典的**“业务逻辑”和“控制逻辑”分离解耦**的编程模式。