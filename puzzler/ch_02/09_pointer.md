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

