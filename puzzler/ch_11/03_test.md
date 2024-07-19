**Go 对测试函数的名称和签名都有哪些规定？**

- 单元测试 Unit，方法名 `Test*` 前缀，参数列表接收 `*testing.T`。
- 基准测试 Benchmark，方法名 `Benchmark*` 前缀，参数列表接收 `*testing.B`。

**go test 主要流程**

1. 准备：检查 `*_test.go` 文件中是否包含测试函数。
2. 编译：生成临时 bin 文件；如果有 dep，也会一同编译。
3. 运行测试函数：如果有 `t.Parallel()` 则会并行执行。
4. 收集测试函数中的断言结果。
5. 生成测试报告。
6. 后处理，删除临时 bin 文件。

```bash
# unit
$ go test
$ go test -v
$ go test -v ./mypackage
$ go test -cover
$ go test -v > result.log

# bench
# -bench=. run all bench tests
# -run=^$  bench only
$ go test -bench=. -run=^$ ./mypackage
```

`*testing.T` 是单元测试中的**结构体**，每个函数都可以接收一个参数，用于记录结果和报告失败。

- **`t.Error(args ...interface{})`**: 记录一个错误信息，但**不会停止测试**。会继续执行其他的测试代码。

- **`t.Errorf(format string, args ...interface{})`**: 记录格式化的错误信息，并继续执行测试。

- **`t.Fail()`**: 标记测试为失败，但**不会停止测试**。测试会继续执行。

- **`t.FailNow()`**: 标记测试为失败，并停止执行当前测试函数。

- **`t.Fatal(args ...interface{})`**: 标记测试为失败，并停止执行当前测试函数。`t.Fatalf` 可以格式化输出信息。

- **`t.Skip(args ...interface{})`**: 跳过当前测试，记录跳过的原因。`t.Skipf` 可以格式化输出信息。

- **`t.Helper()`**: 将调用该方法的函数标记为测试帮助函数，这样它的调用栈信息会被隐藏在测试结果中，方便定位问题。

- **`t.Log`**：记录一条日志信息。

- **`t.Logf`**：记录一条**格式化**日志信息。

`*testing.B` 是基准测试中的**结构体**，用于运行代码的多次迭代，并记录每次迭代所需的时间。

- **`b.N`**: 记录基准测试的迭代次数。`b.N` 是测试框架决定的，用于执行被测试的代码多次以获得稳定的性能测量。
- **`b.ResetTimer()`**: 重置计时器，并从此时开始计算性能。
- **`b.StopTimer()`**: 在基准测试中，有时需要停止计时器，以便进行一些测试前的准备工作，然后再继续计时。
- **`b.StartTimer()`**: 开启计时器。

**单元测试结果**

是否执行成功，（测试包的路径），耗时。

如果测试代码没有发生变动，则直接返回**缓存** cache 结果；缓存路径：`go env GOCACHE`

清理构建 build 缓存：`go clean -cache`

清理测试 test  缓存：`go clean -testcache`

对于失败的测试结果，不会缓存

**基准测试结果**

操作系统，架构，pkg，maxP，迭代执行测试，平均耗时，成功与否，耗时。

