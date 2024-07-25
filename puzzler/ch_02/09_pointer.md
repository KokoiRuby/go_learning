取地址符 `&`，放到一个变量前使用就会返回相应变量的内存地址。

```go
var a int = 10
println("Addr of var a is: ", &a) // 0x...
```

指针变量：**指向的/存放的是一个地址，这个地址指向的空间才是值**。取值符 `*` 获取指针变量保存地址所指向的值。

```go
var ptr *int
ptr = &a
println("value of ptr is: ", ptr)               // 0x...
println("value that ptr points to is: ", *ptr)  // 10
```



![Golang Pointers Illustration](https://d33wubrfki0l68.cloudfront.net/6112acd194b9f21230bda3782a021047d44ced4a/65465/static/5cd5cace553fc78068e470cdc5312ec4/ace37/golang-pointers-illustration.png)



nil 空指针：声明后未赋值 (&v)，值为对应类型的零值。

```go
var ptr *int
println("value of ptr is: ", ptr) // 0
if(ptr == nil)                    // true
```

修改指针变量

```go
num := 1
var ptr *int
ptr = &num
*ptr = 2
println(n) // 2
```

**Go 中的不可寻址的值**（对于不可寻址的，无法使用 `&` 取址操作符）

- 常量（编译时确定，不在内存里）
- 字面量（具体的值，编译直接使用，没有赋值给变量）
  - 对切片字面量的索引结果值虽然也属于临时结果，但却是可寻址的 ∵ 底层有一个数组
- 表达式结果（临时）
- 包级别的未导出标识符（小写）

**引用类型的值**

- 当我们 make 创建一个引用类型时，实际上创建的是一个底层数据结构，返回该类型的值 = 指向这个数据结构的指针。
- 传递和共享数据效率更高，传递时实际上是传递值副本，但由于这个值就是指针，所以也是传递引用。

> uintptr vs. unsafe.Pointer

`uintptr` 是一个整数类型，用于**存储指针的数值**。

```go
uintptr(unsafe.Pointer(&x))
```

`unsafe.Pointer` 是一个特殖的指针类型，**可包含任何类型的指针**，支持类型转换。

```go
unsafe.Pointer(&x)
(*int)unsafe.Pointer(&x)  // convert to *int
*(*int)unsafe.Pointer(&x) // take value of *int
```

