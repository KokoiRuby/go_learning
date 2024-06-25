[toc]

### Overview

[Go](https://go.dev/doc/) 是 Google 开发一种**静态强类型，可快速编译，支持高并发，面向垃圾回收**的语言。`Go = C + Python` 既有 C 静态语言程序运行速度，有能达到 python 动态语言的快速开发。

:confused: **vs?**

- [x] 静态语言：变量的类型在编译时就确定下来，并且在编译时进行类型检查，必须声明变量的类型。C/C++/Java/Golang.
  - :smile: 编译时能更早发现错误，更好运行时性能。
- 动态语言：变量的类型可以在运行时根据赋值语句自动推断，并且在运行时进行类型检查。Python/Ruby.
- [x] 强类型：变量的类型转换必须显式地进行，不会发生隐式的类型转换。JavaScript/PHP.
  - :smile: 可提高代码可读性 & 维护性。
- 弱类型：变量可以被隐式转换为其他类型。

![Static and Dynamic typing? Strong and weak typing? - DEV Community](https://res.cloudinary.com/practicaldev/image/fetch/s--i1yqfSl1--/c_imagga_scale,f_auto,fl_progressive,h_900,q_auto,w_1600/https://miro.medium.com/max/1400/1%2ABddwVWW6hFU0miT9DCbUWQ.png)

### Feature

- 类型安全 - 显式声明
- 内存安全 - 有指针，但是禁指针操作
- 原生并发 - goroutine
- 快速编译 - 静态语言
- 依赖管理 - 包模型
- 内存管理 - 标记/清除



