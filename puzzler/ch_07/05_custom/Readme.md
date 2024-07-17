**自定义包调用**

```bash
$ go mod init custom
```

```bash
$ tree
.
├── go.mod
├── main
├── pkg1
│   └── pkg1.go
└── test_pkg.go
```

```go
// ./pkg1/pkg1.go
package pkg1

var Pack1Int int = 42
var PackFloat = 3.14

func ReturnStr() string {
	return "Hello main!"
}
```

```go
// .test_pkg.go
package main

import (
	"custom/pkg1"   // moduleName/path/to/package
	"fmt"
)

func main() {
	var test1 string
	test1 = pkg1.ReturnStr()
	fmt.Printf("ReturnStr from package1: %s\n", test1)
	fmt.Printf("Integer from package1: %d\n", pkg1.Pack1Int)
	fmt.Printf("Float from package1: %f\n", pkg1.PackFloat)
}
```

**导入外部安装包**

`go install` 用于将包构建并安装到默认的安装目录中。包的源代码和二进制文件将被放置在 `$GOPATH/pkg` 目录下。

`go get` 用于从远程仓库下载并安装包，或者更新已安装的包。

```bash
$ go get github.com/gin-gonic/gin
```

```go
import (
	"github.com/gin-gonic/gin"
)
```

**godoc** 

一个文档生成工具，通过解析项目 `.go` 文件中包含注释的，来生成 HTML 或文本类型的文档。

通过在本地启动一个 web 程序，可以在浏览器来展示项目的文档。

```bash
$ go get golang.org/x/tools/cmd/godoc
$ sudo ln -s $GOPATH/bin/godoc /usr/local/bin/godoc
$ godoc -http=:16060 -play -index
```

`go doc` 在命令行输出 Go 源码相关的文档说明，可查看包 pkg 的相关文档

```bash
$ go doc fmt
$ go doc fmt Printf
```

