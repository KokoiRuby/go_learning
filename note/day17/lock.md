mutex 互斥是常用的同步机制，用于多线程对共享/临界资源的访问控制。确保数据的安全性和一致性。当一个线程访问共享资源时，其他线程不能访问该共享资源，直到该线程访问完毕。

```go
var mutex sync.Mutex
mutex.Lock()   // lock
// TODO
mutex.Unlock() // unlock
```

read/write 读写锁：允许并行读，独占写。适合读多写少，like cache。
```go
var rwMutex sync.RWMutex
rwMutex.Lock()    // w
rwMutex.Unlock()

rwMutex.RLock()   // r
rwMutex.RUnLock()
```

