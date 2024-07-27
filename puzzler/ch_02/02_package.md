### Package

包是**代码组织和模块化**的封装单位：每个程序都由包（通常简称为 pkg）的概念组成，可以使用自身的包或者从其它包中导入内容。

一个包可以由许多以 `.go` 为扩展名的源文件组成。

每一个 `.go` 文件都需要有一个 package 声明（它的归属）。

- 同一个目录下，`.go` 中 package 声明必须一致。
- 同一个目录下， `.go` 和 `_test.go` 文件的 package 声明可以不一致。

每一个 Go 程序必须有一个名为 main 的包，没有 main 就不能执行。

```go
// *.go belongs to package named "main"
package main
// ... 
```

属于同一个包的源文件必须全部被一起编译，**如果对一个包进行更改或重新编译，所有引用了这个包的程序也必须全部重新编译。**

一个 Go 程序是通过 `import` 关键字将多个包**链接**在一起。

```go
import (
   "fmt"
   "os"
)
```

![Introduction to Go packages. What is a Go package? How to use it? | by  Inanc Gumus | Learn Go Programming](https://miro.medium.com/v2/resize:fit:1400/1*16AcelCn5LA1lL7TJPslTw.png)

![Golang: Import Local Packages - Claire Lee - Medium](https://miro.medium.com/v2/resize:fit:1400/1*F2UrGRS_lOKbpQ2HrlxXvg.png)

### [stdlib](https://studygolang.com/pkgdoc)

在 Go 的安装文件里包含了一些**可直接使用**的包 Off-the-shelf → `$GOROOT/src/fmt`

