### client-server

在通常情况下，就是多个客户端（很多请求）对一个（或几个）服务端。

在 Go 中，服务端通常会在一个协程里操作对一个客户端的响应，协程:客户端 = 1:1。

典型做法：客户端请求本身包含了一个 channel，服务端可以用它来发送响应。

```go
type Request struct {
    a, b int;
    replyc chan int;
}
```

```go
type Reply struct { ... }

type Request struct {

    arg1, arg2, arg3 some_type
    replyc chan *Reply

}
```

服务端可以在一个 goroutine 里面为每个请求都分配一个 run() 函数，对请求进行操作并将结果作为响应放到请求内部的通道中。

```go
type binOp func(a, b int) int

func run(op binOp, req *Request) {
    req.replyc <- op(req.a, req.b)
}
```

而服务器通过循环从专门存放请求的通道中不断取出，启动一个协程运行 run() 函数。如果通道为空，那么会阻塞。

```go
func server(op binOp, service chan *Request, quit chan bool) {
	for {
		select {
		case req := <-service:
			go run(op, req)
		case <-quit:
			return
		}
	}
}
```

main 中启动服务器

```go
func startServer(op binOp) (service chan *Request, quit chan bool) {
	service = make(chan *Request)
	quit = make(chan bool)
	go server(op, service, quit)
	return service, quit
}
```

```go
func main() {
	adder, quit := startServer(func(a, b int) int { return a + b })
	const N = 100
	var reqs [N]Request
	for i := 0; i < N; i++ {
		req := &reqs[i]
		req.a = i
		req.b = i + N
		req.replyc = make(chan int)
		adder <- req
	}

	for i := N - 1; i >= 0; i-- { // 顺序无所谓
		if <-reqs[i].replyc != N+2*i {
			fmt.Println("fail at", i)
		} else {
			fmt.Println("Request ", i, "is ok!")
		}
	}
	fmt.Println("done")
	quit <- true
}
```

