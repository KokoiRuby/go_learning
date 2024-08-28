defer：一种**类栈**的延时机制，通常用于**资源释放**。

Go 逐行编译执行到 defer 的时候，会将其语句入栈，暂时不执行。

```go
// file.go
defer file Close()

// db conn
defer connect.close()

// lock
defer mutex.Unlock() 
```

defer 还可用于代码追踪 Tracing：在进入和离开某个函数时打印日志。

```go
func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

func a() {
    trace("a")
    defer untrace("a")
    fmt.Println("in a")
}

func b() {
    trace("b")
    defer untrace("b")
    fmt.Println("in b")
    a()
}

// entering: b
// in b
// entering: a
// in a
// leaving: a
// leaving: b
func main() {
    b()
}
```

### Take-away

1. 函数执行到 return 语句时，会**先执行返回语句，再执行 defer**。
2. 只要**声明函数返回值变量名**，就会在函数**初始化时之赋零值**，而且在函数体作用域可见。
3. 由于先 return 再 defer，**defer 中可修改原本要返回的值**。
4. defer **保证 panic 后依然有效**，确保资源即使遇到 panic 也可以关闭。
5. defer 中包含 panic，仅最后一个可以被 recover 捕获，也就是 **defer 中的 panic 会覆盖 main 中的 panic**。
6. defer 中包含子函数，**入栈时会执行子函数**。

### Exercise

```go
package main

import "fmt"

// return 4
func DeferFunc1(i int) (t int) {  // t = 0
    t = i
    defer func() {
        t += 3
    }()
    return t
}

// return 1
func DeferFunc2(i int) int {
    t := i
    defer func() {
        t += 3
    }()
    return t
}

// return 3
func DeferFunc3(i int) (t int) { // t = 0
    defer func() {
        t += i
    }()
    return 2
}

func DeferFunc4() (t int) {  // t = 0
    defer func(i int) {
        fmt.Println(i)  // 0
        fmt.Println(t)  // 2 
    }(t)
    t = 1
    return 2
}

func main() {
    fmt.Println(DeferFunc1(1)) // 4
    fmt.Println(DeferFunc2(1)) // 1
    fmt.Println(DeferFunc3(1)) // 3
    DeferFunc4() // 0 2
}
```



