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
- [ ] **$GOMAXPROCS** 用于设置应用程序可使用的处理器个数与核数

### [Proxy](https://github.com/goproxy/goproxy.cn/blob/master/README.zh-CN.md)

```bash
$ go env -w GO111MODULE=on
$ go env -w GOPROXY=https://goproxy.cn,direct
```

