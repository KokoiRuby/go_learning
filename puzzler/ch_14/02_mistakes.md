## Mistakes

if 内部使用短变量声明会 shadow 外面的值；**使用 `=`，而不是 `:=`，或者使用不同变量名。**

```go
x := 10
if x := 20; x > 15 {
    fmt.Println("Inside if: x =", x) // 20
}
fmt.Println("Outside if: x =", x)    // 10
```

对字符串频繁操作时，**应使用strings.Builder 或 bytes.Buffer，将字符串内容写入一个缓存中。**。

```go
var b bytes.Buffer
b.WriteString(str)

var b strings.Builder
builder.WriteString("hello")
```

循环内关闭文件的错误姿势，defer 只要在函数返回时才会执行，**在 for 结尾不会执行**。

```go
for _, file := range files {
    if f, err = os.Open(file); err != nil {
        return
    }
    defer f.Close() // wrong, file.go won't close after loop
    f.Process(data)
 }
```

循环外关闭文件的正确测试，使用 defer + 闭包，确保捕获了初始创建的 file 并在返回前能够被关闭

```go
defer func(f *os.File) {
        if err := f.Close(); err != nil {
            log.Printf("Error closing file.go: %v", err)
        }
    }(file)
```

**对于所有值类型使用 `new`，对于所有引用类型 slice/map/channel 使用 `make`**

- `new` 分配类型内存并赋予零值，返回类型指针。

- `make` 分配内存并初始化，然后返回实例的引用（不是指针）。

**不需要将一个指向切片的指针传递给函数**，∵ 切片自身就是一个指针。

```go
func myFunc(sl *[]type) type { ... } // wrong
func myFunc(sl []type) type { ... }
```

**不要使用一个指针指向一个接口类型**。∵ 接口类型本身就是一个指针。

```go
var i myInterface = MyStruct{value: 42}
ptrToInterface := &i // wrong
ptrToInterface := i
```

**使用值类型时误传指针**，值类型看似有内存拷贝，但由于值类型在栈上分配，快速且开销不大；

如果传指针，可能会创建一个对象到堆上，导致额外的内存分配。

```go
type Point struct {
    X, Y int
}

func printPoint(p *Point) {
    fmt.Printf("Point: (%d, %d)\n", p.X, p.Y)
}

p1 := Point{X: 10, Y: 20}
printPoint(&p1)   // wrong
```

**不要使用 bool 进行错误判断，冗余，**直接使用 error 对象非 nil 判断即可。

```go
_, err1 := api.Func1()
if err1 != nil { … }
```

**尽可能以闭包的形式封装你的错误检测**，将错误处理逻辑集中处理，而不是通过多个 if + err 非 nil 判断。

```go
checkErr := func(err error) {
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
}

result, err := divide(10, 0)
checkErr(err)

result, err = divide(10, 2)
checkErr(err)
```

**也可以通过封装，将错误处理和恢复加入到原函数中**

```go
// alias, same signature 
type fType1 func(a int, b string)

func errorHandler(fn fType1) fType1 {
    return func(a int, b string) {
        defer func() {
            if err := recover(); err != nil {
                if e, ok := err.(error); ok {
                    log.Printf("run time panic: %v", e)
                } else {
                    log.Printf("run time panic: %v", err)
                }
            }
        }()
        
        // 执行传入的函数
        fn(a, b)
    }
}

func riskyFunction(a int, b string) {
    if a == 0 {
        panic("a cannot be zero")
    }
    log.Printf("a: %d, b: %s", a, b)
}

safeFunction := errorHandler(riskyFunction)
safeFunction(0, "example")
safeFunction(1, "example")
```

## Pattern

### **comma ok**

1. 函数返回检查错误

```go
_, err := myFunc()
if err != nil {
    // TODO
}
```

2. map 是否包含 key

```go
if v, ok = m[k], ok {
    // TODO
}
```

3. 类型断言

```go
if v, ok := val.(T); ok {
    // TODO
}
```

4. 检查通道是否关闭

```go
for v := range ch {
    // TODO
}

for {
    if v, open := <-ch; !open {
        break
    }
    // TODO
}
```

### defer

1. 关闭文件流 ++闭包

```go
defer func(f *os.File) {
        if err := f.Close(); err != nil {
            // TODO
        }
    }(file)
```

2. 解锁

```go
mutex.Lock()
defer mutex.Unlock()
```

3. 关闭 channel

```go
ch := make(chan type)
defer close(ch)
```

4. recover from panic

```go
defer func() {
    if err := recover(); err != nil {
        log.Printf(“run time panic: %v”, err)
    }
}
```

5. 停止 Ticker

```go
tick := time.NewTicker(updateInterval)
defer tick.Stop()
```

6. 释放进程

```go
p, err := os.StartProcess(…, …, …)
defer p.Release()
```

7. 停止 pprof

```go
pprof.StartCPUProfile(f)
defer pprof.StopCPUProfile()
```

8. 打印页脚 footer

```go
defer printFooter()
```

### Visibility

小写：仅包内可见，适用于不需要导出或仅用于包内部实现的标识符。类私有。

大写：保外可见，适用于需要导出并被其他包使用的标识符。类公有。

### Operator

操作者模式通常用于将一系列操作封装在一个对象中，这些操作可以对另一个对象进行操作。

Go 中**通过接口和结构体实现**。

1. 函数实现运算符：私有函数，通过单独的函数暴露，结合 type-switch

```go
// ungraceful
func addSparseToDense (a *sparseMatrix, b *denseMatrix) *denseMatrix
func addSparseToSparse (a *sparseMatrix, b *sparseMatrix) *sparseMatrix
```

```go
// graceful
func Add(a Matrix, b Matrix) Matrix {
    switch a.(type) {
    case sparseMatrix:
        switch b.(type) {
        case sparseMatrix:
            return addSparseToSparse(a.(sparseMatrix), b.(sparseMatrix))
        case denseMatrix:
            return addSparseToDense(a.(sparseMatrix), b.(denseMatrix))
        …
        default:
            // 不支持的参数
    …
    }
}
```

2. 使用接口和结构体

```go
type Algebraic interface {
    Add(b Algebraic) Algebraic
    Min(b Algebraic) Algebraic
    Mult(b Algebraic) Algebraic
    …
    Elements()
}

type dm denseMatrix {
	//
}

func (a *denseMatrix) Add(b Algebraic) Algebraic {
    switch b.(type) {
    case sparseMatrix:
        return addDenseToSparse(a, b.(sparseMatrix))
    default:
        for x in range b.Elements() …
    …
}
```