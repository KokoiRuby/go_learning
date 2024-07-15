defer：一种**类栈**的延时机制，通常用于**资源释放**。

Go 逐行编译执行到 defer 的时候，会将其语句入栈，暂时不执行。

在函数 return 之前，会先为返回值赋值，然后依次出栈执行 defer。

```go
// file
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

