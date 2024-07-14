### [Install](https://golang.google.cn/doc/install)

### Env

- [x] **$GOROOT** 表示 Go 在你的电脑上的安装位置，它的值一般都是 `$HOME/go`
  - [x] bin：包含可执行文件，如：编译器，Go 工具
  - [ ] doc：包含示例程序，代码工具，本地文档等
  - [ ] lib：包含文档模版
  - [ ] misc：包含与支持 Go 编辑器有关的配置文件以及 cgo 的示例
  - [x] os_arch：包含标准库的包的对象文件（.a）
  - [x] src：包含源代码构建脚本和标准库的包的完整源代码（Go 是一门开源语言）
  - [ ] src/cmd：包含 Go 和 C 的编译器和命令行脚本
- [ ] **$GOARCH** 表示目标机器的处理器架构 - 386 | amd64 | arm
- [ ] **$GOOS** 表示目标机器的操作系统
- [ ] **$GOBIN** 表示编译器和链接器的安装位置
- [x] **$GOPATH** 默认采用和 `$GOROOT` 一样的值 - src | pkg | bin
- [ ] **$GOARM** 专门针对基于 arm 架构的处理器
- [x] **$GOMAXPROCS** 用于设置应用程序可使用的处理器个数与核数

### [Proxy](https://github.com/goproxy/goproxy.cn/blob/master/README.zh-CN.md)

```bash
$ go env -w GO111MODULE=on
$ go env -w GOPROXY=https://goproxy.cn,direct
```

### GOPATH

指定工作空间 workspace 路径，每个工作区存放：

- src 源码
- pkg 存放 `go install` 或 `go build` 编译后的生成的**库文件**（不能直接执行，只能在源码中引用）
- bin 存放 `go install` 或 `go build` 编译后的生成的**可执行文件**

### cmd

自 Go v1.11 开始，引入了 Go Modules 的功能，用于更好地管理项目的依赖关系，而不再完全依赖于 GOPATH。

**依赖查找顺序**：`./go.mod` → `./vendor/` → `$GOPATH` → `$GOROOT`

| cmd           | Desc                                                         |
| ------------- | ------------------------------------------------------------ |
| `go mod init` | 在当前目录中初始化一个新的 Go Module，生成 `go.mod` 用于跟踪项目的依赖关系。 |
| `go mod tidy` | 往 `go.mod` 中添加任何缺少的依赖项，并删除不再需要的依赖项 = 更新依赖。 |
| `go build`    | 在当前目录下编译 Go 代码并生成可执行文件。                   |
| `go install`  | 在当前目录下编译 Go 代码并生成可执行文件，并将生成的可执行文件安装到指定目录中 `$GOPATH/bin`。 |
| `go get`      | 从**远程仓库**获取包并安装到本地。可执行：`$GOPATH/bin`，第三方：`$GOPATH/pkg/mod` |
| `go run`      | 可直接编译并运行 `.go` 文件，但不会生成可执行文件            |

