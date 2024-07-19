**测试**是开发过程的重要部分，也是软件开发生命周期的关键部分。它可以确保应用程序正常运行和满足客户需求。

**[单元测试](https://blog.jetbrains.com/go/2022/11/22/comprehensive-guide-to-testing-in-go/)**是一种测试函数和方法等小段代码的方法。它的用途在于让您及早发现错误，让测试策略更高效，因为它们小且独立，易于维护。

### Simple

Go 中的测试函数以 `Test` 开头，`Test[NameOfFunction]`，并将 [`*testing.T`](https://pkg.go.dev/testing#T) 作为唯一形参。 

命令行 `go test` 可运行测试。（必须先 `go mod init` 初始化包）

### Table-Driven

使用表驱动方式帮助**减少重复**，将测试用例组织为包含输入和所需输出的表，使测试更加**整洁**。

1. 定义输入结构体（切片） = 定义表的列，而表的每一行代表待执行的测试用例
2. 循环执行

```go
tests := []struct {
		name   string
		intput int
		want   string
	}{
		// TODO
	}

for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := Fooer(tt.intput)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
```

### Testing pkg

`testing.T` 类型提供了控制测试执行的方法。

- [`Parallel()`](https://pkg.go.dev/testing#T.Parallel) 并行运行测试

- [`Skip()`](https://pkg.go.dev/testing#T.Skip) 跳过测试

- [`Cleanup()`](https://pkg.go.dev/testing#T.Cleanup) 调用测试清扫

  

**错误 & 日志**

`Log*()` 可用于输出信息。

[`t.Errorf`](https://pkg.go.dev/testing#T.Errorf) 通过在控制台上显示错误消息来指示测试失败。

`t.Error*` 不会停止测试的执行。

`t.Fatal*` 会停止测试的执行。



**并行运行**

[`Parallel()`](https://pkg.go.dev/testing#T.Parallel) 方法则指示测试应并行运行。所有调用此函数的测试都将并行执行。暂停各个调用 `t.Parallel()` 的测试，然后在所有非并行测试完成后将其并行恢复。

`GOMAXPROCS` 环境定义一次可以并行运行多少个测试，这个数字默认等于 CPU 的核数。



**跳过测试**

[`Skip()`](https://pkg.go.dev/testing#T.Skip) 方法可以将单元测试与[集成测试](https://medium.com/@victorsteven/understanding-unit-and-integrationtesting-in-golang-ba60becb778d)分开。集成测试同时验证多个函数和组件，执行起来通常较慢。

`go test` 接受旨在运行 “fast” 测试的 `-test.short` 标志。即运行 `go test -v -test.short` 时测试将被跳过。

需要结合使用 `testing.Short()`（使用 `-short` 时设为 `true`）和 `t.Skip()`。



**清扫**

`Cleanup()` 函数在每个测试（包括子测试）结束时执行

```go
func Test_With_Cleanup(t *testing.T) {
  // Some test code here
    t.Cleanup(func() {
        // cleanup logic
    })
  // more test code here
}
```

[`Helper()`](https://pkg.go.dev/testing#T.Helper) 用途是在测试失败时改进日志。在日志中，helper 函数的行号会被忽略，只报告失败测试的行号，这有助于确定失败的测试。

```go
func helper(t *testing.T) {
  t.Helper()
  // do something
}
func Test_With_Cleanup(t *testing.T) {
  // Some test code here
    helper(t)
  // more test code here
}
```

[`TempDir()`](https://pkg.go.dev/testing#T.TempDir) 是一种自动为测试创建临时目录并在测试完成时删除该文件夹的方法

```go
func TestFooerTempDir(t *testing.T) {
    tmpDir := t.TempDir()
  // your tests
}
```

### [Coverage](https://go.dev/blog/cover#:~:text=Test%20coverage%20is)

覆盖率用于描述运行了多少百分比的测试代码。

`go test -cover` 生成覆盖率报告。

`go test ./... -coverpkg=./...` 可以指定测试的包

`go test -coverprofile=output_filename` 生成覆盖率报告文本文件

```go
// 源码文件:行号范围 执行次数，总执行次数（set 模式下 0 or 1）
mode: set
unit_test/fooer.go:5.30,7.11 2 1
unit_test/fooer.go:7.11,9.3 1 1
unit_test/fooer.go:10.2,10.28 1 1
```

`go tool cover -html=output_filename` 生成覆盖率图形报告

[`-covermode`](https://go.dev/blog/cover#heat-maps)

- `set`：覆盖率基于语句。
- `count`：计数是一条语句运行的次数。它可以显示代码的哪些部分仅被轻微覆盖。
- `atomic`：与计数类似，但用于并行测试。

### Bench

基准测试是一种测试代码性能的方式，通过**多次运行**相同的函数来**验证算法运行时长和内存使用情况**。

- 测试函数需要位于 `*_test` 文件中。
- 函数名称必须以 `Benchmark` 开头。
- 函数必须接受 `*testing.B` 作为唯一形参。
- 测试函数必须包含一个 `for` 循环（以 `b.N` 为其上限）。

```bash
Running tool: /usr/local/go/bin/go test -benchmem -run=^$ -bench ^BenchmarkFooer$ unit_test

goos: linux
goarch: amd64
pkg: unit_test
cpu: 12th Gen Intel(R) Core(TM) i5-1245U
=== RUN   BenchmarkFooer
BenchmarkFooer
# 12 cores 总运行次数，每次耗时，每次操作分配的内存，每次操作额外分配的内存
BenchmarkFooer-12       97705698                12.96 ns/op            5 B/op       0 allocs/op
PASS
ok      unit_test       1.285s
```

### Fuzz

模糊测试使用**随机输入**探索错误或**边缘**用例。 

- 测试函数需要位于 `_test` 文件中。
- 函数名称必须以 `Fuzz` 开头。
- 测试函数必须接受 `testing.F` 作为唯一形参。
- 测试函数必须使用 `f.Add()` 方法定义初始值，即[种子语料库](https://go.dev/security/fuzz/#glos-seed-corpus)。
- 测试函数必须定义[模糊目标](https://go.dev/security/fuzz/#glos-fuzz-target)。

默认情况下，只要没有失败，模糊测试就会无限期地运行。使用 `-fuzztime` 标志指定模糊测试运行的最长时间。