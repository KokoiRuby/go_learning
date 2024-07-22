### :construction_worker: Practice

- 性能分析必须在一个可重复的、稳定的环境中来进行。 
- 必须闲置机器，不能共享硬件，不能做其他曹祖。
- 注意省电模式和过热保护。
- **不要使用虚拟机、共享的云主机**。

### time & %CPU

`time` is a Linux cmd used to **measure the execution time** of a command or program.

- `real`: actual elapsed time from the start to the end of 
- `user`: time elapsed in userspace.
- `sys`: time elapesd in kernel.

:bookmark_tabs: `real` >= `user` + `sys` since ctx sw.

```bash
$ time go run to_run.go
$ /usr/bin/time go run to_run.go  # more details
```

### Golang Mem

```bash
$ go build 05_perf_mem2.go && ./05_perf_mem2
$ top -p $(pidof 05_perf_mem2)
```

**GC Trace** 启用详细的垃圾回收事件日志记录。

```bash
# enable GC tracing
$ GODEBUG='gctrace=1' ./05_perf_mem2
```

```bash
gc 10                                  # GC 次数序号
@0.014s                                # 程序已执行的时间
4%:                                    # GC 占 CPU 百分比。
0.051+1.9+0.002 ms clock,              # GC 标记段、清理和其他相关操作的时间。
1.0+0/1.8/0+0.043 ms cpu,              # 上述期间使用的 CPU 时间
35->35->19 MB,                         # 开始时堆内存/结束时堆内存/实际被程序占用的堆内存
35 MB goal, 0 MB stacks, 0 MB globals, # 全局堆内存，对堆进行操作的内存，对全局变量进行操作内存
20 P                                   # P 数量 = core

gc 22 
@0.128s 
5%: 
0.007+22+0.002 ms clock, 
0.14+0/22/0+0.058 ms cpu, 
516->516->0 MB, 516 MB goal, 0 MB stacks, 0 MB globals, 
20 P

# 强制 GC 已经把 516MB 内存标记为非活跃，全局堆内存下降

gc 23 
@0.150s 
5%: 0.12+0.12+0.010 ms clock, 
2.4+0/0.093/0+0.20 ms cpu, 
0->0->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 
20 P (forced)

gc 24 
@122.175s 
0%: 0.040+0.19+0.005 ms clock, 
0.80+0/0.25/0+0.10 ms cpu, 
0->0->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 
20 P
```

`runtime.ReadMemStats(&runtime.MemStats)` 主线程起一个协程循环执行。

[ms.*](http://golang.org/pkg/runtime/#MemStats)

- `Alloc`：已分配的内存字节数。
- `HeapIdle`：未被使用的堆内存字节数。
- `HeapReleased`：已被返回操作系统的堆内存字节数。**上升说明 GC 将内存归还给了 OS。**

```go
func readMemStats() {
    var ms runtime.MemStats
    runtime.ReadMemStats(&ms)
    log.Printf(" ===> Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes)", 
               ms.Alloc, 
               ms.HeapIdle, 
               ms.HeapReleased)
}

go func() {
    for {
        readMemStats()
        time.Sleep(10 * time.Second)
    }
}()
```

```bash
# [Start].
Alloc:222992(bytes) HeapIdle:3350528(bytes) HeapReleased:3317760(bytes)
# loop begin.
Alloc:154238504(bytes) HeapIdle:272769024(bytes) HeapReleased:149471232(bytes)
# loop end.
Alloc:242968(bytes) HeapIdle:619683840(bytes) HeapReleased:78094336(bytes)
# [force gc].
# [Done].
Alloc:248784(bytes) HeapIdle:619667456(bytes) HeapReleased:80617472(bytes)
```

**pprof** 起一个协程，监听一个本地端口

http://localhost:16060/debug/pprof/heap?debug=1

```go
go func() {
	log.Println(http.ListenAndServe("localhost:16060", nil))
}()
```

```bash
# runtime.MemStats          
Alloc = 360064               # 当前程序已分配的内存字节数，已分配 + 未释放
TotalAlloc = 1503573224      # 程序总共已分配的内存字节数，已分配 + 已释放
Sys = 651677096              # 程序向操作系统申请的内存字节数，堆 + 栈 + 运行时
Lookups = 0                  # Deprecated
Mallocs = 766                # 程序执行的内存分配次数
Frees = 163                  # 程序执行的内存释放次数
HeapAlloc = 360064           # 堆上已分配的内存字节数
HeapSys = 645300224          # 堆向操作系统申请的内存字节数
HeapIdle = 644202496         # 堆中空闲的内存字节数
HeapInuse = 1097728          # 堆中正在使用的内存字节数
HeapReleased = 644161536     # 已经释放给操作系统的堆内存字节数
HeapObjects = 603            # 堆中分配的对象数量
Stack = 622592 / 622592      # 栈内存使用百分比
MSpan = 75360 / 81600        # MSpan 使用百分比
MCache = 24000 / 31200       # MCache 使用百分比
BuckHashSys = 1454672        # 哈希表的内存字节数
GCSys = 2914568              # 垃圾回收器使用的内存字节数
OtherSys = 1272240           # 其他未分类的系统级内存字节数
NextGC = 4194304             # 下一次垃圾回收（GC）的目标堆大小
LastGC = 1721655993505585883 # 上一次垃圾回收的纳秒时间戳
```

### Golang CPU

pprof 主进程监听一个本地端口，协程运行计算任务。

http://localhost:16060/debug/pprof/

profile: cpu, mem, lock, groutine... 默认是 30s 生成一个

```bash
$ go build 05_perf_cpu2.go && ./05_perf_cpu2
```

`go tool` 交互式解析生成的 pprof 文件。

- `flat`：当前函数占用 CPU 的耗时
- `flat%`：:当前函数占用 CPU 的耗时百分比
- `sum%`：函数占用 CPU 的耗时累计百分比
- `cum`：当前函数加上调用当前函数的函数占用 CPU 的总耗时
- `cum%`：当前函数加上调用当前函数的函数占用 CPU 的总耗时百分比

```bash
$ go tool pprof ./05_perf_cpu2 profile
(pprof) top
# sudo apt-get install graphviz
(pprof) web
```

指定生成 profile 的时间，然后直接查看

```bash
$ go tool pprof http://localhost:16060/debug/pprof/profile?seconds=60
# after 60s
$ go tool pprof http://localhost:16060/debug/pprof/profile?seconds=60
```



