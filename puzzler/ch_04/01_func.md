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

Go 函数还可以作为**实参** argument，在其他函数内部调用，一般也称为回调 callback。

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

Go 支持**可变长参数** Variadic，**本质是一个 slice 切片**，可通过 index 访问到各个参数。

使用空接口 `interface{}`，可以接受任何类型的参数，在通过 for-range 和 switch 进行类型判断。

```go
func sum(numbers ...int) int {
	s := 0
	for i := 0; i < len(numbers); i++ {
		s += numbers[i]
	}
	return s
}

func typecheck(..,..,values … interface{}) {
    for _, value := range values {
        switch v := value.(type) {
            case int: …
            case float: …
            case string: …
            case bool: …
            default: …
        }
    }
}
```

Go 函数**不支持重载** Overloading，即同名函数不同参数列表。

原因：简易性，提高可读性，避免同名混淆，避免多余的类型匹配影响性能。

:bookmark_tabs: **vs**

**重载 Overloading**：函数名相同，但参数列表一定不同，返回值可任意。

**重写 Overriding**: **继承**关系中，子类对父类的同名方法重写，参数列表必须相同，作用域修饰符相同或者更大。

### Call by *

Value: 默认值传递，也就是传参的副本；函数执行过程中对副本值的修改不会影响到原来的变量。

Reference: 如果希望函数可以直接修改参数的值，而不是对参数的副本进行操作，你需要将参数的地址传过去，使用 `&`。

:bookmark_tabs: **引用类型：**slice, map, interface, channel。

:construction_worker: 对于指针变量，它保存的值是地址，所以按引用传递也是按值传递。

:construction_worker: 传递指针（64-bit = 8 bytes）的消耗都比传递副本来得少。

:construction_worker: 返回值同理：如果是值类型会拷贝；如果是引用类型，返回的是指针。

```go
// call by ref: a, b points to same addr
func DoSomething(a *A) {
	b = a
}

// call by value: arg a's addr is different (copied), b points to different as well.
func DoSomething(a A) {
    b = &a
}
```