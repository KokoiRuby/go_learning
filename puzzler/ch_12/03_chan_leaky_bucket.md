客户端无限循环执行从某个来源接收的数据，使用一个 Buffer 类型作缓冲读取。

可以保留一个用缓冲通道表示的空闲列表，避免过多分配/释放。该缓冲通道，客户端和服务端共享。

客户端先尝试从通道获取 buffer，如果没有就分配一个新的。当 buffer 都加载完毕后发送给服务端。

```go
var serverChan = make(chan *Buffer)

func client() {
    for {
    	var b *Buffer
    	select {
        case b = <-freeList:
        default:
        	b = new(Buffer)
    }
    loadInto(b)
    serverChan <- b
    }
}
```

服务器端循环接收每一个客户端的消息，处理它，并尝试将 buffer 返回给共享的 buffers 列表

```go
func server() {
    for {
        b := <-serverChan
        process(b)
        select {
            case freeList <- b:
            default:
        }
    }
}
```

