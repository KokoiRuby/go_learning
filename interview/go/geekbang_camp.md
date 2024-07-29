### Basic

> 什么是闭包？有什么好处？有什么缺陷？

闭包是一个函数，该函数绑定了**运行时上下文**，通常是一个匿名函数引用了一个变量，从而构成了一个封闭的环境。理解为在那个时刻作了一个**快照**。

:smile: 逻辑封装延迟执行；减少全局变量的使用

:cry: 被闭包引用的变量无法被释放，可能存在内存泄露

```go
// var from inside
func closure1() func() {
	i := 0
	return func() {
		i++
        println(i)
	}
}

// var from outside
func closure2(name string) func() string {
	return func() string {
		return "Hello, " + name
	}
}
```

> 什么情况下会出现栈溢出？

程序使用的栈空间超过了系统为其分配的大小限制 `ulimit -s`。

- 递归函数没有递归出口，即无限递归
- 大量局部变量 (分配在栈上)
- 无线循环中执行函数
- 过多 goroutine (2K)

> 什么是不定参数？调用方法的时候，不定参数可以传入 0 个值吗？方法内部怎么使用不定参数？

不定参数即 varridic 不定长参数 `...type`；可以传 0 个参数。本质上是一个 []type **切片**，可通过 for-range 遍历。

> 什么是 defer ？你能解释一下 defer 的运作机制吗？

defer 是**类栈 LIFO** 的延迟机制，保证函数退出前一定执行，通常用于**资源释放/追踪/修改命名返回值**。

> 一个方法内部 defer 能不能超过 8 个？

Go 没有限定 defer 数量限制，要尽量控制不超过 8，以让编译器实现内联 inline 提高性能。

> defer 内部能不能修改返回值？怎么改？

可以，前提是返回值需命名。

顺序：初始化命名返回值 → return 语句赋值返回值 → defer 中对返回值进行修改

```go
func deferReturn() (result int) {
	result = 1
	defer func() {
		result *= 2
	}()
	return result
}
```

> defer 实现机制

**堆分配**：`_defer` 结构体，被 GC 管理。

**栈分配**：goroutine 栈，**速度更快**，栈不被 GC 管理。

**开发编码 Open Code**：**内联**：将 `defer` 调用**内联 inline** 到包含 `defer` 的函数中，以**提高性能**和减少函数调用的开销。**条件**：

- defer 数量 <= 8
- defer 不能在循环中执行
- return * defer <= 15

> 数组和切片有什么区别？

类型不同：数组值类型；切片引用类型

长度：数组定长；切片不定长且可扩容

切片底层对应一个数组。

> 切片怎么扩容？

容量不够时，Go 会创建一个新的底层数组，再进行元素搬迁，**策略**：

1. <= 256 (之前 1024)，扩 2 倍 
2. \> 256，扩 1.25 倍

**翻倍内存对于小切片容易产生浪费，降低阈值能减少内存的分配和释放操作。**

原则：扩多了浪费内存；扩少了频繁执行内存分配操作。

> 为什么没有缩容？

泛型工具库，实现切片删除辅助功能，++ 缩容机制。

### Gin

> 什么是 Gin 的 middleware？能用来解决什么问题？

中间件 AKA plugin/handler/filter/interceptor → **AOP 解决方案：所有业务都关联通用逻辑处理进行抽离（形成一个切面）**，比如：AuthN & AuthZ, Logging, Monitoring。

**Request → []middleware → Business Logic**

具体案例支撑：针对 IP 限流，针对用户等级进行分级服务...

> 什么是跨域问题，怎么解决？

跨域问题：由浏览器的**同源策略（Same-Origin Policy）**引起的，该策略要求浏览器只允许页面加载来自同一域的资源，以**防止恶意网站获取用户的敏感信息**。

不同域：protocol/domain/port 任一一个不同就判定不同域。

解决：CORS 跨域资源共享 Cross-Origin Resource Sharing。使用 Gin 开源插件。

> 跨域问题需要设置哪些头部？

`Access-Control-[Allow-Origin|Allow-Headers|Allow-Metods]`

> Plus: 最好提一下自己的自研 Gin 插件库