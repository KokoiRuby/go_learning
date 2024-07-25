`go module` **依赖管理工具** since v1.11，production since v1.14，兼容 `$GOPATH`；

## GOPATH

所有代码都存放在固定 `$GOPATH/src` 下，`go get` 下载的远端依赖会安装在该目录。

- `/bin`：存储所编译生成的二进制文件。
- `/pkg`：存储预编译的目标文件，以加快程序的后续编译速度。
- `/src`：存储所有 `.go` 文件或源码。

```go
$ go env | grep GOPATH
```

:cry: 无版本控制：`go get` 无法携带版本信息 → 项目依赖库版本难以保持一致。

## Go Module

允许在 `$GOPATH/src` 外任何目录下适用创建项目。

```bash
# switch
$ go env -w GO111MODULE=on
```

模块 Module 是包 Package 的**集合**。

`go.mod` 声明了模块的名称及其依赖项，`go.sum` 记录 checksum 确保完整性。



<img src="https://images.ctfassets.net/ro61k101ee59/2jEvZ29rRWjEbiU4guNdAP/ab59d827ad145553c59dac5049a3dfe1/go-module.png?w=978&q=75" alt="Migrating our monorepo seamlessly from Dep to Go Modules" style="zoom: 67%;" />

### cmd

| cmd               | desc                                 |
| ----------------- | ------------------------------------ |
| `go mod init`     | 生成 `go.mod` 文件。                 |
| `go mod download` | 下载 `go.mod` 文件中指明的所有依赖。 |
| `go mod tidy`     | 整理现有的依赖。                     |
| `go mod graph`    | 查看现有的依赖结构。                 |
| `go mod edit`     | 编辑 `go.mod` 文件。                 |
| `go mod vendor`   | 导出项目所有的依赖到 vendor 目录。   |
| `go mod verify`   | 校验一个模块是否被篡改过。           |
| `go mod why`      | 查看为什么需要依赖某模块。           |

### Quickstart

```bash
# init
$ mkdir backend && cd backend
$ go mod init backend
$ cat go.mod
```

```go
// main.go
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}
```

```bash
# tidy to update dep
$ go mod tidy

# chk again
# require → dep & ver.
$ cat go.mod
```

```bash
# list dep
$ go list -m all

# update
$ go get -u github.com/gin-gonic/gin

# chk hist ver
$ go list -m -versions github.com/gin-gonic/gin

# get ver
$ go get github.com/gin-gonic/gin@vx.y.z

# or modify go.mod
$ go mod edit -require="github.com/gin-gonic/gin@v1.1.4"
$ go mod tidy
```

### Deepin

**语义导入版本 Semantic Import Versioning**

vX.Y.Z where X is major 主版本, Y is minor 次版本, Z is patch

- X 不同达标不兼容
- X 相同，Y 大的后向兼容
- Z 不影响兼容性

**如果同一个包的新旧版本是兼容的，那么它们的包导入路径应该是相同的。**

```go
import {
	"github.com/sirupsen/logrus" // works for both v1.7.0 & v1.8.1
}
```

对于不兼容的版本，需要将版本号引入

```go
import (
    "github.com/sirupsen/logrus"           // v1.7.0 & v1.8.1
    logv2 "github.com/sirupsen/logrus/v2"  // v2.0.4
)
```

**最小版本选择 Minimal Version Selection**

**Go 会在该项目依赖项的所有版本中，选出符合项目整体要求的“最小版本”**

proj → A v1.5.0 → C v1.1.0 where latest C v1.7.0

​     → B v1.2.0 → **C v1.3.0**