### OOP

OOP 是一种基于 “对象” 概念的编程范式。**类是对象的封装，set(属性+方法)；对象是类的实例。**

**封装 Encapsulation**：隐藏对象的内部属性和实现细节，仅对外提供公有方法进行调用。Go 通过**首字母大小写**决定是否能在包外访问。

```go
type Person struct { name string }

// public constructor
func NewPerson() *Person { return &Person{} }

// public setter/getter
func (p *Person) SetName(name string) { p.name = name }
func (p *Person) GetName() string { return p.name }
```

**继承 Inheritance**：子类继承父类的属性和方法。**Go 通过匿名字段组合的方式实现继承。**

```go
type Parent struct { Name string }

type Child_1 struct {
    Parent
    Feature_1 string
}

type Child_2 struct {
    Parent
    Feature_2 string
}
```

多态 Polymorphism：父类指针指向具体子类的实例，并实现不同子类的方法从而表现出不同的行为。**Go 通过接口实现多态。**

```go
type Action interface { play() }

func (c1 *Child_1) play() { ... }
func (c2 *Child_2) play() { ... }

var p Parent
p := &Child_1{ ... }
p.play()
p := &Child_2{ ... }
p.play()
```

> Golang 属于 OOP 语言嘛？

Go 有类型和方法，允许面向对象的编程风格，但没有类型层次，即强父子类联系。

Go 方法更为通用，可为任何类型定义，不局限于结构体（类）。

> vs. Java

Go 不允许函数重载，函数签名唯一。

Go 并发性能优于 Java。

Go 通过匿名成员组合实现继承，支持多继承。

> vs. Python

Go 静态语言（编译时决定变量类型）；Python 动态语言（运行时决定变量类型）。

Go 有内置并发实现；Python 则没有。

Go & Python 都属于强类型，即必须要进行显式类型转换。

Go 允许开发人员管理内存；Python 则完全由 Python VM 管理，不对开发人员开放。

### Basics

> Pros

基于协程原生支持高并发。

语法简单，上手快。

C (快速编译) + Python (快速开发)。

类型安全，显式声明

内存安全，有指针但禁指针操作。

基于 module & package 的依赖管理。

GC

> make vs. new

共同点：根据变量类型分配内存。**Go 编译器的逃逸分析**，即决定变量是分配到堆上还是栈上。

**若函数外对该变量无引用，则优先放到栈中**。

**若函数外对该变量有引用，则必定放到堆中**。

不同点：

`new` 接收一个类型，**返回**一个指向该类型内存地址的**指针**，同时赋 zero-val。

`make` 只用于 slice/map/chan 内存创建，**返回**的就是这 3 **类型本身**，而不是指向它们的指针。

> defer

defer 是一种**类栈 LIFO**的延时机制，通常用于**资源释放**，比如文件/数据库/连接关闭，以及异常捕获，代码追踪。

底层：多 `_defer` 实例构成的**单链表**，每次插入于链头，执行时从链头取出。

1. 函数执行到 return 语句时，会**先执行返回语句，再执行 defer**。
2. 只要**声明函数返回值变量名**，就会在函数**初始化时之赋零值**，而且在函数体作用域可见。
3. 由于先 return 再 defer，**defer 中可修改原本要返回的值**。
4. defer **保证 panic 后依然有效**，确保资源即使遇到 panic 也可以关闭。
5. defer 中包含 panic，仅最后一个可以被 recover 捕获，也就是 **defer 中的 panic 会覆盖 main 中的 panic**。
6. defer 中包含子函数，**入栈时会执行子函数**。

> uint8 溢出

最大存储 255 (11111111)，++ 溢出变成 0 (1|00000000)，即最高位被丢弃。

> rune

rune alias to int32, byte alias to int8.

Go 源码都按 Unicode 编码规范中 UTF-8 编码格式进行编码。

每一个字符对应一个独特 unicode 码点，即 rune。每一个 rune 由若干字节构成。

string 既可以表示为 `[]rune` 也可以表示 `[]byte`。

> '' vs. "" vs. ``

'' 单引号用于表示 rune 字面值。

"" 双引号用于表示 string 字面值。

`` 反引号用于表示原始字符串字面值，不进行转义。

> tag resolv & reflection

Go 中解析的结构体字段 tag 是通过反射实现的。

```go
type Person struct {
	Name  string `key:"value"`
}
nameTag := reflect.TypeOf(dog).Field(0).Tag.Get("key")
```

正常情况下，源码编译后变量转换成内存地址，变量名不会被编译器写入到可执行部分，运行中程序无法获取自身的信息。

反射 Reflection 指在程序**运行期**对程序本身进行**访问和修改**的能力。

基于反射，能够将 `interface{}` 类型变量转换成 `Type` & `Value` 类型。

- `Kind` 指一个变量的基本类型分类，静态。
- `Type` 指一个变量的具体类型信息，动态。

```go
// Kind: int
// Type: int
var a int = 42

// Kind: ptr
// Type: *int
var b *int = &a

// Kind: slice
// Type: []int
var c []int = []int{1, 2, 3}

// Kind: map
// Type: map[string]int
var d map[string]int = map[string]int{"one": 1, "two": 2}
```

相互转换

```go
var v interface{}

t    := reflect.TypeOf(v)  // interface{} → reflect.Type
rVal := reflect.ValueOf(v) // interface{} → reflect.Value
rVal.Kind()                // reflect.Value.Kind()
rVal.Type()                // reflect.Value.Type()

iVal := rVal.interface()   // reflect.Value → interface{}
orig := iVal.(type)        // interface{} → type
```

修改不可变类型值

```go
str := "str"
rfv := reflect.ValueOf(&str)
rfv.Elem().SetString("STR")
fmt.Println(str) // STR
```

调用结构体方法

```go
t := reflect.TypeOf(i)
v := reflect.ValueOf(i)

if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
    // without args
    val.Method(x).Call(nil)
    
    // with args
    var params []reflect.Value
    params = append(params, reflect.ValueOf(10), reflect.ValueOf(20))
    res := val.Method(x).Call(params)
}
```

> 调用函数传入结构体，传值还是引用？

Go 函数参数传递都是值传递，拷贝一份数据到栈内存，改变不影响原本数据。

vs. 传引用，即传指针/地址，修改会影响到原本数据。

> goroutine 什么情况下会阻塞？

1. 原子/互斥/通道操作。
2. IO → Go NetPoller 基于 Linux epoll 实现 IO 多路复用。
3. 系统调用阻塞 M，调度器引入其它 M 来重新进行绑定 P 并执行后续的 G。
4. 执行 sleep，阻塞 M。

> select 底层数据结构 or 项目中如何使用 select？

select 提供了一种 IO 多路复用机制，每一个分支可以处理一个管道，通常用于轮询监听，在某个条件下 break。

- 每个 case 语句仅能处理一个管道
- 多个 case 语句的执行顺序是随机的
- 存在 default 语句，select 将不会阻塞，会立即执行该分支并一直运行；若其他通道长时间阻塞，会不断轮询消耗 CPU

> panic

- 数组/切片越界
- 空指针调用
- 过早关闭 HTTP 响应体
- 除以 0
- 向已经关闭的 chan 发送
- 重复关闭 chan
- 关闭未初始化的 chan
- 未初始化 map，即没有 make
- waitGroup 计数为负数
- 类型断言不匹配，建议 `s, ok := a.(string)`

> 如何实现 while

```go
// pre
for {
    if cond {
        break
    }
    // TODO
    // post
}
```

> 如何实现 set - Add/Remove/Contains

set 是一个集合，其本质就是一个 list，只是元素不能重复。

内置一个 map 判断 key 是否重复，value 使用 struct{} 节约内存。

3pp: `golang-set`

> 如何复用一个接口的方法？

基于接口嵌套，作匿名成员**组合**进

```go
type BasicInterface interface {
	SayHello() string
}

type AdvancedInterface interface {
	BasicInterface
	SayGoodbye() string
}
```

> _ 

1. 忽略返回值占位符 `if _, ok := func()`
2. 接口断言，判断类型是否实现了某接口类型 `var _ Interface = Type`
3. 导入包仅初始化，不适用包中的功能 `import _ "/path/to/package"`

> goroutine 创建的时候若要传一个参数进去有什么要注意的点？

闭包捕获：goroutine 在创建时捕获了变量的引用，而不是变量的值。

```go
for i := 0; i < 5; i++ {
	go func() {
		fmt.Println(i)
	}()
}
```

要避免这种问题，可以通过局部变量作为参数传递给闭包来捕获变量的值。

```go
for i := 0; i < 5; i++ {
	i := i 
	go func(i int) {
		fmt.Println(i)
	}(i)
}
```

> 单测

需创建 `*_test.go` 后缀的文件。命名规范 `package *_test`。

测试方法名 `Test_` ，[`*testing.T`](https://pkg.go.dev/testing#T) 作为唯一形参。

> delve

A **debugger** for the Golang.



> go module

依赖管理工具，允许在 `$GOPATH/src` 外任何目录下适用创建项目，代替 `$GOPATH`，并引入了版本管理。

模块 Module 是包 Package 的**集合**。

`go.mod` 声明了模块的名称及其依赖项，`go.sum` 记录 checksum 确保完整性。

> 浅拷贝 vs. 深拷贝

浅拷贝：拷贝引用/地址/指针，修改会影响原值；深拷贝：拷贝值，修改不会影响原值。

> Go 多返回值如何实现的？

Go 编译器处理多返回值时，会在函数栈中为每个返回值分配独立的空间。这意味着每个返回值都有自己独立的内存位置，并且函数调用者可以直接访问这些返回值。

> Go 不同类型如何判等？

对于 slice/map/struct 用 `reflect.DeepEqual`，其余 `==` 即可。

> Go init

一个包下可以有多个 init()，一个 .go 也可以有多个 init()；在 main() 前执行；通常用于初始化。

顺序：常量 → 全局变量 → 从被导入的最深层包开始 init()，层层递出到最外层 main()。

> uintptr vs. unsafe.Pointer

`uintptr` 是一种整数类型，用于**存储指针的数值**。

`unsafe.Pointer` 是一个特殊的指针类型，用于执行**任意类型指针之间的转换**。

> 什么是 goroutine？

定义：goroutine 是 Go 中的一种用户级线程，由 Go 运行时管理。

使用：`go` 关键字 + 函数。

优势：轻量 2KB，创建销毁切换开销极小。

调度：Go runtime 自行维护。

通信：同步原语 ([RW]mutex/Cond/Atomic/WaitGroup/Once/Channel)

应用：高并发（网络/计算）

### Slice

> vs. array

相同：连续内存保存同类型数据；支持下标随机访问；len() & cap() 分别获取长度即元素个数 & 容量。

不同：数组定长，切片自动扩容（migrate，地址改变）；数组值类型，切片引用类型；

:construction_worker: [x]int, [y]int 是不同类型，∵ 数组长度也属于类型的一部分。

> 底层数据结构

\*array 指向底层数组的地址 + 切片长度/容量 len/cap。

> 扩容

创建新的底层数组并 migrate → if len < 1024，cap *= 2; else len >= 1024，cap *= 1.25

> array/slice 作参数区别

数组还是切片，在函数中传递的时候若没有指定为指针传递的话，都是值传递。

但切片属于引用类型，传递的值是一个地址/指针，修改会影响源数据。

若不想修改源数据，可以通过 `copy(dst, src []Type)` 复制一个切片指向一个新的底层数组，再传参。

> 从数组中取一个相同大小的切片有成本吗？

```go
// 成本非常低，只是建了一个包含指针、长度和容量的切片结构体
arr := [5]int{1, 2, 3, 4, 5}
sl := arr[:]
```

> 从切片中取一个相同大小的数组有成本吗？

```go
sl := []int{1, 2, 3, 4, 5}
var arr [5]int
// 成本高，O(n) 逐个复制到新数组中
copy(arr[:], sl)
```

### Map

> 使用注意点

K 必须可判等；slice/map/func 不可作 K。

struct 作 K 必须所有字段都可判等，浮点数作字段可能存在精度问题，同时要提供 Key() & Hash() 方法。

K 无序，有序可选用 3pp。

使用前一定要通过 make 进行初始化否则 panic。

非并发安全，并发需使用 `sync.Map`。

> 删除一个 K，内存会释放嘛？

删除 `map` 中的键会标记相关内存为可回收，不会立即释放，后续由 GC 管理。

> nil map vs. 空 map

nil map 即未 make，取出得 nil；赋值/删除K 都会 panic。

空 map 是 make 但未赋值。

> 底层

hash table 一个哈希表里可以有多个 bucket，每个 bucket 就保存了 map 中的一个或一组键值对。

```go
type hmap struct {     
   count     int                  // 元素个数     
   flags     uint8     
   B         uint8                // 扩容常量相关字段 B 是 buckets 数组的长度的对数 2^B     
   noverflow uint16               // 溢出的 bucket 个数     
   hash0     uint32               // hash seed     
   buckets    unsafe.Pointer      // buckets 数组指针     
   oldbuckets unsafe.Pointer      // 结构扩容的时候用于赋值的 buckets 数组     
   nevacuate  uintptr             // 搬迁进度     
   extra *mapextra                // 用于扩容的指针 
}
```

data 连续存放是为了节约内存避免 padding 带来的浪费。

```go
type bmap struct {
   tophash [8]uint8 // 存储哈希值的高 8 位
   data    byte[1]  // key value 数据: key/key/key/.../value/value/value...
   overflow *bmap   // 溢出 bucket 的地址
}
```

![](https://cdn.nlark.com/yuque/0/2022/png/22219483/1661789834784-c60b0cb4-96be-4c4c-8978-2bfc9ca716b9.png#averageHue=%23414141&clientId=uef4c3b7a-0bed-4&errorMessage=unknown%20error&from=paste&id=ueb411ead&originHeight=260&originWidth=664&originalType=url&ratio=1&rotation=0&showTitle=false&size=23558&status=error&style=none&taskId=u3f259a38-81e8-46dc-912d-aee35205a60&title=#averageHue=%23414141&errorMessage=unknown%20error&id=WzvZN&originHeight=260&originWidth=664&originalType=binary&ratio=1&rotation=0&showTitle=false&status=error&style=none)

> 哈希冲突

当有两个或以上的 K 被哈希到了同一个 bucket 时即发生冲突。Go 采用链地址法解决冲突。

每个 bucket 存放 8 个 K/V，超过则新建一个 bucket，原 bucket 通过 overflow 指向溢出新建的 bucket，以类链表的方式连接起来。

![img](https://www.topgoer.cn/uploads/gozhuanjia/images/m_a8b9e5919d9951a71c1c36445dd68521_r.png)

> 负载因子 & 扩容

load factor = # of K / # of Bucket。过小，说明空间利用率低；过大，说明冲突严重，存取效率低。

扩容条件：

- lf > 6.5 时，也即平均每个 bucket 存储的键值对达到 6.5 个。
- overflow > 2^15 时，也即 overflow 数量超过 32768 时。

**增量扩容**：满足条件时，开辟原 2 倍的 bucket 数组空间。oldbuckets 指向原数组；buckets 指向新数组。

**逐步搬迁**：新 KV 进 buckets；每次访问 map，从旧数组中取部分 KV 到新数组中。



![img](https://www.topgoer.cn/uploads/gozhuanjia/images/m_2f0122f26e5d66ca91e6820ace6b379b_r.png)

> 查找

1. 根据 key 值算出哈希值
2. 取哈希值低位与 hmap.B 取模确定 bucket 位置
3. 取哈希值高位在 tophash 数组中查询
4. 若 tophash[i] 中存储值也哈希值相等，则去找到该 bucket 中的 key 值进行比较
5. 当前 bucket 没有找到，则继续从下个 overflow 的 bucket 中查找。
6. 若当前处于搬迁过程，则优先从 oldbuckets 查找

注：若查找不到，也不会返回空值，而是返回相应类型的零值。

> 插入

1. 根据 key 值算出哈希值
2. 取哈希值低位与 hmap.B 取模确定 bucket 位置
3. 查找该 key 是否已经存在，若存在则直接更新值
4. 若没找到将 key，将 key 插入

> 可对 map 中得一个元素取址嘛？

不能，map 动态扩容会导致原指针失效。

### sync.Map

Go stdlib 中支持并发安全的 map，适用于读多写少。

1. **读写分离**：read & dirty，Load() 优先访问 read，未命中则访问 dirty。
2. **延迟写入**：Store() 先查找 read 是否存在 K，若存在且标记“未删除” 则写 read，否则 dirty。
3. **读原子写加锁**：read 通过 atomic.Value 保存数据无需加锁；dirty 通过互斥锁保护写。
4. **读写转换**：misses 计数器记录未命中 read，当达到阈值，dirty 会**转换**成 read，同时 dirty 置 nil。
5. **条目淘汰**：先查找 read 是否有 K，若有则标记未已删除。下次转换时从 read 中删除。

### Interface

Duck Type 鸭子类型：动态语言的风格，即一个对象有效的语义，不是由继承自特定的类或实现特定的接口，而**由其方法属性结合决定**。

**Go 通过接口实现鸭子类型。**

> 方法

方法是开发者为自定义类型添加的行为。

vs. 函数

- 声明不同：方法多一个接收者。
- 调用不同：方法通过类型的**实例调用**。函数可以直接通过**函数名调用**。
- 声明位置：方法必须声明在同一个包内；函数可以声明在任何包中。

> 值类型接收者 vs. 引用类型接收者

若实现了接收者是值类型的方法，会隐含地也实现了接收者是指针类型的方法。反之不成立。

若接收者是值类型，无论调用者是对象还是对象指针，修改的都是对象的副本，不影响调用者

若接收者是引用类型，则调用者修改的是指针指向的对象本身。

> iface vs. eface

iface 用来在运行时表示拥有方法的接口类型的变量。tab 即动态类型，data 即动态值。

接口零值是指动态类型和动态值都为 nil。当仅且当这两部分的值都为 nil 的情况下，这个接口值就才会被认为接口值 == nil。

```go
type iface struct {
    tab  *itab               // 指向一个 itab 实体，表示接口的类型以及赋给这个接口的实体类型
    data unsafe.Pointer      // "指向"当前被赋值给接口类型变量的具体类型变量值
}

type itab struct {
    inter  *interfacetype  // 描述了接口的类型
    _type  *_type          // 实体的类型，包括内存对齐方式，大小等
    link   *itab
    hash   uint32 
    bad    bool   
    inhash bool   
    unused [2]byte
    fun    [1]uintptr      // 存储的是第一个方法的函数指针
}

type interfacetype struct {
    typ     _type      // 各种数据类型的结构体
    pkgpath name       // 定义了接口的包名
    mhdr    []imethod  // 接口所定义的函数列表
}
```

![go-interface-iface.png](https://blog.frognew.com/images/2018/11/go-interface-iface.png)

eface 在运行时表示方法集为空的接口类型变量：`interface{}`

```go
type eface struct {
    _type *_type          // 空接口所承载的具体的实体类型
    data  unsafe.Pointer  // 具体的值的指针
}
```

![go-interface-eface.png](https://blog.frognew.com/images/2018/11/go-interface-eface.png)

> **自动检测类型是否实现接口？**

利用编译器检测和 `_` 下划线占位符实现

```go
// *myWriter 类型是否实现了 io.Writer 接口
var _ io.Writer = (*myWriter)(nil)
// myWriter 类型是否实现了 io.Writer 接口
var _ io.Writer = myWriter{}
```

> 类型转换 vs. 类型断言

类型转换、类型断言本质都是把**一个类型转换成另外一个**。不同之处在于，**类型断言是对接口变量进行的操作，通常结合 switch 判断执行对应 case**。

```go
var i srcType = val
var f dstType
f = dstType(i)
fmt.Printf("%T, %v\n", f, f)
```

```go
type Student struct {
    Name string
}

var i interface{} = &Student{}
_, ok := i.(Student)
```

### Context

上下文 [`context.Context`](https://draveness.me/golang/tree/context.Context) 用于在多个 goroutine 之间**传递信号/请求、超时和截止日期等信息**。

- `Deadline`：返回 [`context.Context`](https://draveness.me/golang/tree/context.Context) 完成工作的截止日期；
- `Done`：返回一个 chan，该 chan 会在当前工作**完成后**/上下文被**取消后关闭**，多次调用 `Done` 方法会返回同一个 chan；
- `Err`：返回 [`context.Context`](https://draveness.me/golang/tree/context.Context) 结束的原因，它只会在 `Done` 方法对应的 Channel 关闭时返回非空的值；
  - 若 ctx 取消，返回 `Canceled` 错误；
  - 若 ctx 超时，返回 `DeadlineExceeded` 错误；
- `Value`：从 [`context.Context`](https://draveness.me/golang/tree/context.Context) 中获取键对应的值，对同一个上下文，多次调用并传入相同的 `Key` 会返回相同的结果。

```go
type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}
```

> UC

1. **取消信号**：取消其他 goroutine。例：在处理 HTTP 请求时，若客户端关闭了连接，可以使用 `context` 取消所有相关的后台操作。

```go
cancelCtx, cancel := context.WithCancel(context.Background())
go longRunningTask(cancelCtx)  // pass cancel ctx to goroutine
time.Sleep(3 * time.Second)
cancel()                       // cancel from main
```

2. **超时控制**：超时自动发送取消信号。例：在数据库查询或网络请求时，可以使用 `context` 设置一个超时时间，以防止长时间的等待。

- `context.WithDeadline()` 派生一个带绝对截至时间的 ctxt。
- `context.WithTimeout()` 派生一个带持续时间的 ctxt。

```go
parentCtx := context.Background()
deadline := time.Now().Add(5 * time.Second)
ctx, cancel := context.WithDeadline(parentCtx, deadline)
```

```go
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handle(ctx, 500*time.Millisecond)
    // go handle(ctx, 1500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}
```

3. **传递数据**：goroutine 之间传递请求范围的数据；例：meta-like 请求的唯一 ID、用户认证信息。

```go
withValCtx := context.WithValue(context.Background(), "K", "V")
value := ctx.Value("K");
```

### Channel

> 通道并发安全嘛？

并发安全，内置互斥锁。

**不要通过共享内存来通信 (Lock)，而是通过通信来共享内存 (Channel)。**

通道适用于多 goroutine 相互之间传递数据

> 数据结构

`buf` 保存 goroutine 之间传递数据的循环链表。

`send/recvx` 记录此循环链表当前发送或接收数据的下标值。

`sendq/recvq` 保存向该 chan 发送和从 chan 接收数据的 goroutine 队列。

`lock` 保证 channel 写入和读取数据时线程安全的锁

```go
type hchan struct {
    //channel 分为无缓冲和有缓冲两种。
    //对于有缓冲的 channel 存储数据，借助的是如下循环数组的结构
	qcount   uint           // 循环数组中的元素数量
	dataqsiz uint           // 循环数组的长度
	buf      unsafe.Pointer // 指向底层循环链表的指针
	elemsize uint16         //能够收发元素的大小
  

	closed   uint32   //channel 是否关闭的标志
	elemtype *_type   //channel 中的元素类型
  
    //有缓冲 channel 内的缓冲数组会被作为一个“环型”来使用。
    //当下标超过数组容量后会回到第一个位置，所以需要有两个字段记录当前读和写的下标位置
	sendx    uint   // 下一次发送数据的下标位置
	recvx    uint   // 下一次读取数据的下标位置
   
    //当循环数组中没有数据时，收到了接收请求，那么接收数据的变量地址将会写入读等待队列
    //当循环数组中数据已满时，收到了发送请求，那么发送数据的变量地址将写入写等待队列
	recvq    waitq  // 读等待队列
	sendq    waitq  // 写等待队列

	lock mutex //互斥锁，保证读写channel时不存在并发竞争问题
}
```

> 对于 nil、关闭的 channel、有数据的 channel，再进行读写/关闭会如何？

**“空读写阻塞，写关闭异常，读关闭空零”**

- nil 读写永久阻塞
- 写关闭 panic
- 读关闭返回零值

> send/recv/close 流程

**发送：**

1. 若 recvq 不为空，说明 buf 中没有数据或者没有 buf，此时直接从 recvq 取出 G,并把数据写入，最后把该 G 唤醒，结束发送过程。
2. 若 buf 中有空余位置，将数据写入缓冲区，结束发送过程。
3. 若 buf 中没有空余位置，将待发送数据写入 G，将当前 G 加入 sendq，进入睡眠，等待被读 goroutine 唤醒。

**接收：**

1. 若 sendq不为空，且没有 buf，直接从 sendq 中取出 G，把 G 中数据读出，最后把 G 唤醒，结束读取过程；
2. 若 sendq 不为空，此时说明 buf 已满，从缓冲区中首部读出数据，把G中数据写入缓冲区尾部，把 G 唤醒，结束读取过程；
3. 若 buf 中有数据，则从缓冲区取出数据，结束读取过程；
4. 将当前 G 加入 recvq，进入睡眠，等待被写 G 唤s醒；

**关闭：**把 recvq 中的 G 全部唤醒，本该写入 G 的数据位置为 nil。把 sendq 中的 G 全部唤醒，但这些 G 会 panic。

> unbuffered vs. buffered

- Unbufferred：无缓冲，即没有指定 cap；发送和接收是同步的。
  - 发送会立刻阻塞管道，直到接收。
  - 接收会立刻阻塞管道，直到发送。
  - UC：立即阻塞，强制同步协作。
- Buffered：缓冲，指定 cap；管道未满/未空不会阻塞。
  - 缓冲满则阻塞发送
  - 缓冲空会阻塞接收
  - UC：（不同速率）生产消费。
- Send-only
  - UC：限制生产者只能生产，传递给函数时限制其只能发送。
- Receive-only
  - UC：限制消费者只能消费，传递给函数时限制其只能接收。

### GMP

1. **概述**：GMP 模型是 Go 语言运行时用于管理 goroutine 的调度模型。
2. **G (Goroutine)**：Go 语言中的轻量级线程。每个 goroutine 包含需要执行的函数、栈空间以及其他调度信息。
3. **M (Machine)**：代表 Machine，是操作系统的内核线程，负责执行 goroutine，多个 M 可以同时运行在多个 CPU 核上。
4. **P (Processor)**：代表 Processor，是一个逻辑处理器，每个 P 维护一个本地队列，存储要运行的 goroutine。
   - P 的数量由 `GOMAXPROCS` 环境变量设置，默认值为 CPU 核心数。
5. **GMP 模型的工作原理**：
   - 当一个 goroutine 被创建时，它会被放入到某个 P 的本地队列中。
   - P 从本地队列中取出 goroutine 并分配给 M 执行。
   - 如果一个 P 的本地队列为空，会从其他 P 的队列中窃取任务，保证工作负载均衡。
   - 当一个 goroutine 发生阻塞时，M 会释放 P，去寻找其他可运行的 goroutine，避免资源浪费。
6. **优势**：支持高并发，最大程度利用 CPU。

> Process vs. Thread vs. Coruotine

进程：运行的程序实例，系统资源分配的单元，每个进程拥有独立的 vmem。

线程：进程内执行实例，CPU 调度/时间片分配的单元，多线程共享线程资源。

协程：用户空间轻量级线程，协程调度器负责将协程和线程绑定。

> **为什么要有 P？**

P 本地队列减少全局队列依赖以及全局锁竞争。Work Stealing 从全局队列以及其他 P 本地队列窃取 G 运行减少空转，提高 CPU 利用率。

**为什么不直接在 M 上加本地队列呢？**

- 一般 M 远多于 P，M 阻塞且不够的情况会创建新的 M
- M 增加，本地队列也会增加，同时会增加 Work Stealing 的开销

> 调度器设计策略

**复用线程**：避免频繁的创建/销毁，实现线程复用。

- work stealing：当 M 无可运行的 G 时，尝试从其他线程绑定的 P 偷取 G，而不是销毁 M。
- hand off：当 M 因为 G 进行系统调用阻塞时，M 释放绑定的 P，把 P 转移给其他空闲的 M 执行。

**并行**：支持同时 GOMAXPROCS 数量的 goroutine 运行。

**抢占**：一个 goroutine 最多占用 CPU 10ms，防止其他 goroutine 被饿死

**全局队列**：当 M 执行 work stealin g从其他 P 偷不到 G 时，它可以从全局队列获取 G。

> 抢占式 vs. 信号量调度

1. **抢占式调度**：
   - **定义**：操作系统可以在任何时候暂停当前执行的任务，并将处理器分配给另一个任务。这种调度策略允许操作系统在任务执行期间中断任务，通常是基于任务的优先级或时间片大小来决定是否进行切换。
   - **特点**：抢占式调度使得操作系统可以更灵活地管理任务的执行，可以确保高优先级的任务及时执行而不被低优先级的任务阻塞。这种调度方式通常用于实时操作系统或需要强调任务响应时间的系统中。
   - **示例**：常见的抢占式调度策略包括优先级调度（Priority Scheduling）和时间片轮转调度（Round Robin Scheduling）。
2. **信号量调度**：
   - **定义**：任务根据信号量的状态来决定是否可以继续执行。信号量通常用于控制对共享资源的访问，通过对信号量进行操作，任务可以进入临界区或等待共享资源的释放。
   - **特点**：信号量调度通过信号量的机制来协调并发任务的执行顺序，确保对共享资源的访问是线程安全的。信号量也可以用于避免死锁和资源竞争问题。
   - **示例**：常见的信号量包括互斥信号量（Mutex Semaphore）和条件变量（Condition Variable），它们通常与锁机制一起使用，用于控制对共享资源的访问。

### Lock

> 除了 mutex 以外还有那些方式安全读写共享变量？

1. 一个 goroutine 将共享变量放入全局 channel 中，其它 goroutine 通过 channel 进行读写操作，操作完再放入。
2. 基于 sync.WaitGroup 使用个数为 1 的信号量（semaphore）实现互斥。

> Go 如何实现原子操作？

原子操作就是不可中断的操作，在某个值的原子操作执行的过程中，CPU 绝对不会再去执行其他针对该值的操作，即绝不会被中断。

`sync.Atomic` 支持 add, cas (compare & swap), load, store, swap 操作以及类型：int32、 int64、uint32、uint64、uintptr，unsafe.Pointer。

**传地址**，∵ 原子操作基于内存。

`atomic.value` 容器，以原子方式存储和加载**任意类型**的值。

- 不能存 nil，也不能时 interface{}，否则 panic
- 存储的第一个值的类型，决定了之后存储值得类型，否则 panic
- 不要存引用类型

> 悲观锁、乐观锁是什么？Mutex 是悲观锁还是乐观锁？

**悲观**：数据修改前上锁，结束后解锁。适用于冲突发生概率高，读少写多。

**乐观**：数据修改后检测是否有冲突再决定上锁，适合冲突发生概率低，读多写少。

Go 属于悲观锁。

> Mutex 模式

**正常**：

1. 当前 mutex 只有一个 G 来获取，那么没有竞争，直接返回。
2. 新 G 进来，如果当前 mutex 已经被获取了，则该 G 进入一个 waiter 队列，在 mutex 被释放后，waiter 出队列。于队列中 G 会自旋。
3. 新 G 进来，mutex 处于空闲状态，将参与竞争。新 G 具有优势，进 waiter 队列 G 可能会迟迟出不了队列就会被饿死。若 waiter 队列获取锁时间超过 1ms，那么就就进入饥饿模式。

**饥饿**：

mutex 持有则直接把锁交给 waiter 队列队头的 G，新 G 不会尝试获取锁也不会自旋，会进入等待队列。回退正常模式：

1. waiter 队列中已无 G。
2. waiter 队列等待锁小于 1ms。

> goroutine 的自旋占用资源如何解决？

自旋锁是指当一个线程在获取锁的时候，如果锁已经被其他线程获取，那么该线程将循环等待，不断尝试获取锁 → 不断消耗 CPU。

自旋条件

1. 自旋次数未超过 4 次
2. 多核处理器
3. GOMAXPROCS > 1
4. P 本地队列为空

mutex 会让当前的 goroutine 去空转，在空转完后再次调用原子 CAS 去尝试性的占有锁资源，直到不满足自旋条件，则最终会加入到全局等待队列里。

### Concurrency

> Go 中主协程如何等待其余协程退出?

`sync.WaitGroup` 类信号量等待一组协程结束。

主协程 Add() 添加协程计数可一次多加，子协程结束后 Done() 减去一个计数，主协程 Wait() 阻塞直到所有的任务完成。

> 如何控制并发数？

1. 带缓冲通道，利用满写阻塞，空读阻塞特性。
2. 3pp 协程池：[ants](https://github.com/panjf2000/ants), [tunny](https://github.com/Jeffail/tunny), [go-commons-pool](https://github.com/skoo87/go-commons-pool)

> 多个 goroutine 对同一个 map 写会 panic，异常是否可以用 defer 捕获？

可以捕获异常，但是只能捕获一次。

> 如何优雅的实现一个 goroutine 池

[TODO](https://blog.csdn.net/finghting321/article/details/106492915/)

请求数大于消费能力怎么设计协程池？

### GC

> Go 如何实现 GC？

- v1.3 之前的所谓**标记清除**是先启动 STW 暂停，然后执行标记，再执行数据回收，最后停止 STW。
- v1.3 标记清除做了点优化，流程是：先启动 STW 暂停，然后执行标记，停止 STW，最后再执行数据回收。
- v1.5 **三色标记**主要是插入屏障和删除屏障，写入屏障的流程：程序开始，全部标记为白色。
  1. 所有的对象放到白色集合
  2. 遍历一次根节点，得到灰色节点
  3. 遍历灰色节点，将可达的对象，从白色标记灰色，遍历之后的灰色标记成黑色
  4. 由于并发特性，此刻外界向在堆中的对象发生添加对象，以及在栈中的对象添加对象，在堆中的对象会触发插入屏障机制，栈中的对象不触发
  5. 由于堆中对象插入屏障，则会把堆中黑色对象添加的白色对象改成灰色，栈中的黑色对象添加的白色对象依然是白色
  6. 循环第 5 步，直到没有灰色节点
  7. 在准备回收白色前，重新遍历扫描一次栈空间，加上 STW 暂停保护栈，防止外界干扰（有新的白色会被添加成黑色）在 STW 中，将栈中的对象一次三色标记，直到没有灰色
  8. 停止 STW，清除白色。至于删除写屏障，则是遍历灰色节点的时候出现可达的节点被删除，这个时候触发删除写屏障，这个可达的被删除的节点也是灰色，等循环三色标记之后，直到没有灰色节点，然后清理白色，删除写屏障会造成一个对象即使被删除了最后一个指向它的指针也依旧可以活过这一轮，在下一轮 GC 中被清理掉
- v1.8 **混合写屏障**。
  1. GC 开始将栈上的对象全部扫描并标记为黑色(之后不再进行第二次重复扫描，无需 STW
  2. GC 期间，任何在栈上创建的新对象，均为黑色
  3. 被删除的对象标记为灰色
  4. 被添加的对象标记为灰色。

> GC 触发时机

**系统触发**：

1. gcTriggerHeap：当所分配的堆大小达到阈值（由控制器计算的触发堆的大小）时，将会触发。
2. gcTriggerTime：当距离上一个 GC 周期的时间超过一定时间时，将会触发。时间周期以runtime.forcegcperiod 变量为准，默认 2 分钟
3. gcTriggerCycle：如果没有开启 GC，则启动 GC。

**手动触发**：`runtime.GC`

> GC 中 STW 时机

1. 在开始新的一轮 GC 周期前，需要调用 gcWaitOnMark 方法上一轮 GC 的标记结束（含扫描终止、标记、或标记终止等）。
2. 开始新的一轮 GC 周期，调用 gcStart 方法触发 GC 行为，开始扫描标记阶段。
3. 需要调用 gcWaitOnMark 方法等待，直到当前 GC 周期的扫描、标记、标记终止完成。
4. 需要调用 sweepone 方法，扫描未扫除的堆跨度，并持续扫除，保证清理完成。在等待扫除完毕前的阻塞时间，会调用 Gosched 让出。
5. 在本轮 GC 已经基本完成后，会调用 mProf_PostSweep 方法。以此记录最后一次标记终止时的堆配置文件快照。
6. 结束，释放 M。

### Memory

> 什么情况下内存会泄露？

1. G 在执行时被阻塞而无法退出，虽然一个 G 只有 2KB，但数量上来后的消耗也不能忽视。
2. 互斥锁未释放 or 死锁（一直等在那边嘛）
3. time.Ticker 每隔指定的时间就会向通道内写数据。作为循环触发器，必须调用 stop 方法才会停止，否则一直占用内存
4. 字符串的截取引发临时性的内存泄漏，子串依旧保留底层数组的一部分，无法被释放。可将截取的子串转换为新的字符串，从而确保底层字节数组得以释放。

```go
s1 := "This is a very large string..."
s2 := s1[0:10]         // nok
s1 = ""
fmt.Println(s2)

s2 := string(s1[0:10]) // ok
```

5. 切片截取引起子切片内存泄漏，同理，只保留子切片会导致原始切片和底层数组无法被回收。

```go
var s0 = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
s1 := s0[:3]
so = nil
```

6. 函数数组传参引发内存泄漏。∵ 数组属于值类型，传参会拷贝；若数组足够大，那么会消耗大量内存

> 怎么定位排查内存泄漏问题？

- `GODEBUG='gctrace=1' ./after/go/build`
- `runtime.ReadMemStats(&runtime.MemStats)` 主线程起一个协程循环执行
- pprof 起一个协程，监听一个本地端口

> 逃逸

逃逸：本该分配到栈上的变量，跑到了堆上，这就导致了内存逃逸。

- 如果函数外部没有引用，则优先放到栈中；
- 如果函数外部存在引用，则必定放到堆中；

影响：栈内存函数调用结束后回收，分配回收开销小；一旦逃逸到堆上，就只能 GC

情况：

1. 方法内返回局部变量指针。
2. 向 channel 发送指针数据。
3. 在闭包中引用包外的值。
4. 在 slice 或 map 中存储指针。
5. 切片（扩容后）长度太大。
6. 在 interface 类型上调用方法。

> 哪些对象分配在堆上，哪些对象分配在栈上？Channel 分配在栈上还是堆上？

**栈上分配**：

- 基本数据类型（int、bool、float64 等）和结构体（struct）的实例通常会分配在栈上。分配速度快。
- 函数的参数和局部变量也会分配在调用栈上。**若函数局部变量作为返回被外部引用时，比如局部变量的指针返回，那么编译时就会逃逸到堆上。**

**堆上分配**：

- 使用 `new` 关键字创建的变量会在堆上分配。使用 `make` 创建的 slice、map 和 channel 会在堆上分配。分配速度慢，易产生内存碎片。
- 复杂的数据结构，比如包含引用类型的结构体或数组，通常会在堆上分配。

> 介绍一下大对象小对象，为什么小对象多了会造成 GC 压力？

小对象 <= 32k，其余都是大对象。小对象通过 mspan 分配内存；大对象则直接由 mheap 分配内存。

小对象频繁分配和释放可能会导致堆内存的碎片化，进而申请不够时触发 GC。

小对象声明周期短，更容易被标记从而触发 GC。

小对象过多会增加扫描的开销。

### Compile

> GOROOT vs. GOPATH

`GOROOT` 指向安装的 Go 语言的根目录的路径。

`GOPATH` 指定你的工作目录（Workspace）的路径

- `src` 目录包含你的 Go 源代码文件。
- `pkg` 目录包含编译后的包对象文件。
- `bin` 目录包含可执行文件。

> 编译过程

Go 程序并不能直接运行，每条句必须转化为一系列的低级机器语言指令，将这些指令打包到一起，并以二进制可执行文件存储。

![](https://golang.design/go-questions/compile/assets/7.png#id=sXo2v&originHeight=396&originWidth=1920&originalType=binary&ratio=1&rotation=0&showTitle=false&status=done&style=none&title=#errorMessage=unknown%20error&id=zD9xO&originHeight=396&originWidth=1920&originalType=binary&ratio=1&rotation=0&showTitle=false&status=error&style=none)

> 编译指令

| cmd           | Desc                                                         |
| ------------- | ------------------------------------------------------------ |
| `go mod init` | 在当前目录中初始化一个新的 Go Module，生成 `go.mod` 用于跟踪项目的依赖关系。 |
| `go mod tidy` | 往 `go.mod` 中添加任何缺少的依赖项，并删除不再需要的依赖项 = 更新依赖。 |
| `go build`    | 在当前目录下编译 Go 代码并生成可执行文件。                   |
| `go install`  | 在当前目录下编译 Go 代码并生成可执行文件，并将生成的可执行文件安装到指定目录中 `$GOPATH/bin`。 |
| `go get`      | 从**远程仓库**获取包并安装到本地。可执行：`$GOPATH/bin`，第三方：`$GOPATH/pkg/mod` |
| `go run`      | 可直接编译并运行 `.go` 文件，但不会生成可执行文件            |