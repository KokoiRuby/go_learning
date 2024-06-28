**函数**是基本的代码块，用于执行一个任务。实现封装，使用调用。

至少包含一个名为 main 的函数，作为该程序的入口。

```go
func func_name ( [param list] ) [return_types] {
   // body
}
```

- `func` 函数声明关键字
- `func_name` 函数名
- `param list` 形参列表 + `return_types` 返回值，可多返回（使用 `_` 忽略返回） → **函数签名 Signature**



Go 函数本身也是一种**数据类型**，**可以赋值给一个变量**；该变量就是函数类型的变量，通过该变量实现对函数类型的调用。

类型 `%T` 就是函数的签名

```go
func getSum(a, b int) int {
	return a + b
}

x := getSum // %T → func(int, int) int
```

Go 函数还可以作为**实参** argument

```go
func myFunc(funvar func(int, int) int, a, b int) int {
	return funvar(a, b)
}

b := myFunc(getSum, 1, 2)
```

Go 支持**(多)返回值命名**

```go
func cal(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return sum, sub
}
```

Go 支持**可变长参数** Variadic，**本质是一个 slice 切片**，可通过 index 访问到各个参数

```go
func sum(numbers ...int) int {
	s := 0
	for i := 0; i < len(numbers); i++ {
		s += numbers[i]
	}
	return s
}
```

