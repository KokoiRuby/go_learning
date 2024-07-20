**工厂模式**：不将通道作为参数传递给协程，而用函数来生成一个通道并返回（工厂角色）。函数内有个匿名函数被协程调用。

``` go
func newChan chan type {
    ch := make(chan type)
    go func() {
       ch <- ... 
    }()
    return ch
}

func main() {
    ch := newChan()
    go proc(ch)
}
```



