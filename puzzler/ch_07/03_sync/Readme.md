## [Sync](https://pkg.go.dev/sync)

> 临界资源 critical 被多个线程共享 → 竞争条件 Race Condition，进而破坏共享数据的一致性。

同步的目的：**避免**多线程同一时刻对同一资源进行操作；**协调**它们（应该以什么样的方式有序访问）。

Go **同步原语** sync primitives：

- `Mutex`：用于保护共享资源，以防止多个 goroutine 同时访问。
- `RWMutex`：允许多个读操作并发进行，但在写操作时会阻塞所有读和写操作。
- `Cond`：条件变量提供了一种机制，使得 goroutine 可以等待某些条件发生，并在条件发生时被通知。
- `Atomic`：提供原子操作，能完全屏蔽竞争条件且无法被中断。
- `WaitGroup`：用于等待一组 goroutine 完成工作。
- `Once`：确保某个操作只执行一次，适用于需要保证初始化只发生一次的场景。
- **`Channel`：用于 goroutine 之间的通信和同步，能够安全地传递数据和信号。**

### [Mutex](https://pkg.go.dev/sync#Mutex)

:construction_worker: 不要重复锁定互斥锁，否则会触发**死锁**；∵ 是**不可重入锁**，同一个 goroutine 持有情况下不能再次获取。

:construction_worker: 不要忘记解锁互斥锁，务必在上锁时使用 `defer` 确保最后解锁。

:construction_worker: 不要对尚未锁定或者已解锁的互斥锁解锁，否则 panic。

:construction_worker: 不要在多个函数之间直接传递互斥锁（传递副本），易死锁，不易维护。

```go
var mutex sync.Mutex

mutex.Lock()
// TODO
defer mutex.UnLock()
```

### [RWMutex](https://pkg.go.dev/sync#RWMutex)

读写锁是 Mutex 的一种扩展。

:construction_worker: 不能同时写；不能同时读写；只能同时读。

:construction_worker: 解锁未上锁，锁上未解锁都会触发 panic。

```go
var rwmutex sync.RWMutex

// read lock
rwmutex.RLock()
// TODO
defer rwmutex.RUnLock()

// write lock
rwmutex.Lock()
// TODO
defer rwmutex.UnLock()
```

### [Cond](https://pkg.go.dev/sync#Cond)

**条件变量**并不具备独立的锁机制，因此它需要一个已存在的互斥锁来进行操作。

不是被用来保护临界资源，主要用于**协调**。当共享资源发生变化时，它可被用来**通知**被互斥锁阻塞的线程。

- `Wait()` 使当前 goroutine **等待条件满足/等待对方通知**，需要先持有一个互斥锁。
- `Signal()` 唤醒**队头**等待条件变量的 goroutine。
- `Broadcast()` 唤醒**所有**等待条件变量的 goroutine。

```go
var mailbox uint8     // 0 empty, 1 full
var lock sync.RWMutex
sendCond := sync.NewCond(&lock)          // W lock as cond
recvCond := sync.NewCond(lock.RLocker()) // R lock as cond 
```

```go
// sender
lock.Lock()
for mailbox == 1 {   // if full
    sendCond.Wait()
}
mailbox = 1
lock.Unlock()
recvCond.Signal()  // wake up receiver
```

```go
// receiver
lock.RLock()
for mailbox == 0 {  // if empty
    recvCond.Wait()
}
mailbox = 0
lock.RUnlock()
sendCond.Signal()   // wake up sender
```

**:confused: 为什么要先锁才能调用 `Wait()` 呢？**

Wait() 主要做了四件事：

1. 将调用它的 goroutine 加入当前 cond 的**通知队列（队尾）**。
2. 解锁 cond 基于的互斥锁。
3. 阻塞 goroutine 使其等待通知
4. 通知来了唤醒并重新上锁。

∴ 提前上锁是为了 Wait() 种解锁，否则 panic。如果在未解锁的情况下阻塞，那么谁来解锁呢？

**:confused: 为什么使用 for？**

可能存在虚假唤醒，条件可能尚未满足，应该再次调用 wait 方法，∴ 需要反复检查。

### [Atomic](https://pkg.go.dev/sync/atomic)

Go 运行时为了确保公平会频繁进行 goroutine 的切换：运行时 ←→ 非运行时。

虽然互斥锁保证 goroutine 执行临界代码时不被干扰，但它仍可能被中断，进而破坏了操作的原子性。

**原子操作确保操作不会被中断。**

原子操作要求快速足够简单，否则会影响执行效率（毕竟不能被中断）。

**支持操作**：add, cas (compare & swap), load, store, swap

**支持数据类型**：int32、 int64、uint32、uint64、uintptr，unsafe.Pointer

```go
var counter int32 = 100
atomic.AddInt32(&counter, 1)
atomic.AddInt32(&counter, -1)

var value int32 = 42
if atomic.CompareAndSwapInt32(&value, 42, 100) {
    fmt.Println("Value was 42, changed to 100")
}

var value int32 = 42
loadedValue := atomic.LoadInt32(&value)

var value int32
atomic.StoreInt32(&value, 42)

var value int32 = 42
oldValue := atomic.SwapInt32(&value, 100)
```

**:confused: 为什么传入的是地址？**原子操作是**基于内存操作**，传入内存地址才可以进行精准操作。

**:confused: 读还需要原子嘛？**确保写完了才能读，类读写锁。

**:confused: cas vs. swap？**cas 由多个操作构成：比较/判等/交换，+for 可实现自旋锁，类互斥锁。

```go
for {
    if atomic.CompareAndSwapInt32(&num2, 10, 0) {
        fmt.Println("The second number has gone to zero.")
        break
    }
    time.Sleep(time.Millisecond * 500)
}
```

`atomic.value` 容器，以原子方式存储和加载**任意类型**的值。

属于结构体，赋值/传递会拷贝，修改不会改变原有的。

:construction_worker: ​不能存 nil，也不能时 interface{}，否则 panic。

:construction_worker: 存储的第一个值的类型，决定了之后存储值得类型，否则 panic。​

:construction_worker: 不要存引用类型​

```go
var value atomic.Value
value.Store(42)
// value.Store("42") // nok
value.Load()
value.Store("42")
// value.Store(42) // nok
value.Load()
```

```go
var box atomic.Value
v := []int{1, 2, 3}
box.Store(v)
v[1] = 4 // unsafe

store := func(v []int) {
    replica := make([]int, len(v))
    copy(replica, v)
    box.Store(replica) 
}
store(v)
v[2] = 5 // safe
```

### WaitGroup

并发安全，类计数器结构体。

**“统一 Add，再并发 Done，最后 Wait”**

- `Add()` 在主 goroutine 记录需要等待的 goroutine 的数量。
- `Done()` 在需要等待的 goroutine 中 defer 来减少计数器。
- `Wait()` 阻塞主 goroutine 等待所有 goroutine 完成任务再继续进行。

**:confused: 计数可以小于 0 嘛？** ダメ！否则 panic。要避免 `Add()` & `Wait()` 在不同 goroutine 中并发。

### Once

确保某个操作**只执行一次**，即使在多个 goroutine 中并发调用。

主要用于初始化操作，这些操作只应该被执行一次，以避免多次初始化带来的开销或错误。

`Do()` 接收一个**无参数无返回值的函数**，并确保这个函数只会被执行一次。

内置 done unit32 字段：记录函数被调用的次数，首次调用 0 → 1；因为要求**原子**操作，所以类型受限。

## [Context](https://pkg.go.dev/context)

用于在多个 goroutine 之间传递取消信号、超时和截止日期等信息 → 生命周期管理 LCM。

**创建**

`context.Background()` 空上下文，作为根节点。

`context.TODO()` 空上下文，但未决定如何使用。

```go

```
