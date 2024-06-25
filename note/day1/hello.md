### Hello World

Go 语言的基本构成：

- 包声明 package
- 引入包 import
- 函数 func
- 变量 var
- 语句 & 表达式 sentence & expor
- 注释 comment

```go
// *.go belongs to which package.
package main

// package to import, so we could use func inside package.
// $GOROOT/src/fmt
import "fmt"

// "main" func as entrypoint, run first if not init()
func main() {
	fmt.Println("Hello world")
}
```

```bash
# run
$ go run hello.go

# build bin & run
$ go build
$ ./hello
```

:warning: **go 编译器是逐行编译的，因此一行只写一条语句**，不要把多条语句写在同一行，否则报错。

:warning: **go 定义的变量或者 import 包如果没有使用到，代码不能编译通过，避免冗余**。

### escape

| Char |  Desc   |
| :--: | :-----: |
| `\t` |   tab   |
| `\n` | newline |
| `\\` |    \    |
| `\r` |  enter  |
| `"`  |    "    |