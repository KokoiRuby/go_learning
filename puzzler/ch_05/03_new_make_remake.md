`new `和 `make` 都是内建的函数，用于在堆 heap 上分配类型内存。

### var

通过 `var` 关键字变量声明不赋值的默认类型 zero-val 零值。

对于**引用类型**，比如指针类型，不仅要声明它，**还有为其分配内存空间**，否则值存到哪里呢？

对于**值类型**，Go runtime 默认已经帮我们分配好了。

```go
var i *int
*i = 10    // panic: runtime error: invalid mem addr or nil ptr derefer
```

`new` 接收一个类型，**返回**一个指向该类型内存地址的**指针**，同时赋 zero-val。

```go
var i *int
i = new(int)  // malloc
*i = 10       // ok
```

`make` 只能用于 slice/map/chan 内存创建，**返回**的就是这 3 **类型本身**，而不是指向它们的指针。

∵ 它们本身就是引用类型，所以没必要返回它们的指针。



<img src="https://zh.mojotv.cn/assets/image/golang_make_new.webp" alt="Go教程:17-make和new的区别| 🐶❤️🦀" style="zoom:50%;" />