package main

import "fmt"

func foo1(x *int) func() {
	return func() {
		*x = *x + 1
		fmt.Printf("foo1 val = %d\n", *x)
	}
}

func foo2(x int) func() {
	return func() {
		x = x + 1
		fmt.Printf("foo2 val = %d\n", x)
	}
}

func main() {
	x := 133
	f1 := foo1(&x) // f1, f2 都是保存了 x=133 闭包，其中 f1 传递引用（修改会影响外部）；f2 传递值(修改不会影响外部)
	f2 := foo2(x)
	f1() // 134
	f2() // 134
	f1() // 135
	f2() // 135

	x = 233 // x 是闭包外的变量
	f1()    // 234, f1 传递引用，外部变量变化会影响到闭包
	f2()    // 136, f2 传递值，外部变量变化不会影响到闭包，仍然是上一个 135
	f1()    // 235
	f2()    // 137

	// x = 235
	foo1(&x)() // 236，新的闭包
	foo2(x)()  // 237，新的闭包
	foo1(&x)() // 237
	// x = 237
	foo2(x)() // 238，新的闭包
	foo2(x)() // 238，新的闭包，不相关

}
