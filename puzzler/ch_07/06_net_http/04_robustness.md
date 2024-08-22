一个 web 服务必须具备健壮性，不能因 panic 就直接崩溃。

一个方法是可以在每一个 handler 中去使用 `defer/recover`，但这会有大量重复代码。

优雅的方式：**采用 decorator + 闭包处理封装 Handlefunc**

```go
type HandleFunc func(http.ResponseWriter, *http.Request)

// decorators
func logPanics(hf HandleFunc) HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[%v] caught panic: %v", r.RemoteAddr, err)
			}
		}()
		hf(w, r)
	}
}

http.Handle("/test1", http.HandlerFunc(logPanics(SimpleServer)))
http.Handle("/test2", http.HandlerFunc(logPanics(FuncServer)))
```

