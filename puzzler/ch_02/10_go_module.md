`go module` since v1.11 增加的依赖管理工具，v1.13 默认开启，兼容 `$GOPATH`；

允许 `$GOPATH/src` 外任何目录下适用创建项目。

模块 Module 是一个包 Package 的集合。

`go.mod` 声明了模块的名称及其依赖项，`go.sum` 记录 checksum 确保完整性。



![Migrating our monorepo seamlessly from Dep to Go Modules](https://images.ctfassets.net/ro61k101ee59/2jEvZ29rRWjEbiU4guNdAP/ab59d827ad145553c59dac5049a3dfe1/go-module.png?w=978&q=75)



```bash
# enable
$ export GO111MODULE=on

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

