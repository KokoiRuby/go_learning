进程：程序在操作系统中的一次执行过程，运行中的程序，是系统进行**资源分配**的基本单位。

线程：进程的一个执行实例，能够**接收 CPU 时间片调度**的最小单元。

![Process vs Thread: What's the Difference? - javatpoint](https://static.javatpoint.com/difference/images/process-vs-thread3.png)

:confused: **vs?**

- Parallelism 并发：多线程同时运行在多核。
- Concurrency 并行：多线程竞争单核。

只要线程数 > CPU core#，就会同时出现并行和并发。

![img](https://substackcdn.com/image/fetch/w_1456,c_limit,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2F20fe44e5-5085-418c-b039-b72ad863a9a2_1200x1000.png)

### Process

**PCB Process Control Block 进程控制块**：描述进程的当前状态。

- `Pointer`：指向下一个 PCB
- `Process state`：运行状态
- `Process ID`：唯一标识
- `Program counter`：进程下一条指令的地址
- `Priority`：优先级
- `Registers`：寄存执行过程中产生的数据
- `List of open files`：由进程打开的文件列表

:bookmark_tabs: PCB 保存在主存中，OS 使用**链表**来维护所有进程的 PCB

- 动态插入/删除 = 进程创建/中止 O(1)
- 遍历 O(N)

![process-control-block](https://cdn1.byjus.com/wp-content/uploads/2022/06/process-control-block.png)

**进程状态**

- `New` ：进程创建
- `Ready` ：进程加载进主存并等待执行
- `Run`：接收调度并分配 CPU 时间片
- `Terminate`：进程执行完毕；同时 OS 删除 PCB。
- `Block/Wait`：IO 阻塞，结束后进入 Ready 状态（等待再次调度）
- `Suspend Ready/Wait`：高优先级抢占

![img](https://cdn1.byjus.com/wp-content/uploads/2022/08/word-image-15.png)

**上下文切换 Ctxt switch** 实现多进程共享 CPU 资源

1. 保存当前进程 ctx 进 PCB。
2. 将当前进程置入队列 Ready/Wait。
3. 新进程接收调度
4. 加载新进程状态并开始执行，执行完毕后再次保存 ctx 到 PCB。
5. 重新从之前的进程加载 PCB 中 ctx 并恢复执行。

![Context-Switching-in-OS](https://cdn1.byjus.com/wp-content/uploads/2022/08/context-switching-in-os.png)

### Thread

**TCB Thread Control Block 线程控制块**：描述线程的当前状态。

**多线程共享进程资源 vmem 地址空间，切换的开销要小很多，且相互通信无需内核中断可直接访问。**

```c
typedef struct {
    int thread_id;            // 线程标识符
    int state;                // 线程状态（运行、就绪、阻塞等）
    void *stack_pointer;      // 栈指针
    void *program_counter;    // 程序计数器
    RegisterSet registers;    // 寄存器状态
    int priority;             // 线程优先级
    SyncInfo sync_info;       // 同步信息
    // 其他线程特定数据
} ThreadControlBlock;
```

**线程状态**

- `Ready`：就绪等待 CPU 分配
- `Running`：运行中
- `Blocked`：等待
- `Terminated`：执行完毕 or 强制中止

:bookmark_tabs: 线程 TCB 也是**链表**形式相互串联，依附在 PCB 上。

![Thread control blocks](https://people.cs.rutgers.edu/~pxk/416/notes/images/05-tcb.png)

:bookmark_tabs: 每个线程在 stack 栈上有**独立**的内存空间。

![img](https://velog.velcdn.com/images%2Fxogh20321%2Fpost%2Fe140c821-eb3b-4924-8ef3-863bfc53b45f%2FUntitled.png)