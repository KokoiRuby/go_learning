[net/http](net/http) provides HTTP client and server implementations.

`http.Request` 描述了客户端请求，内含一个 `URL` 属性。

`http.URL` 描述了 web 服务器的地址，内含 url 字符串的 `Path` 属性。

监听并启动 HTTP Server。

```go
err := http.ListenAndServe(":8080", nil)
err := http.ListenAndServe("localhost:8080", nil)
```

解析表单 Form & Query param in URL。

```go
// http://localhost:8080/submit?name=John&age=30
// or
// curl -X POST -d "name=John&age=30" http://localhost:8080/submit
func formHandler1(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	age := req.FormValue("age")
}

// or
func formHandler2(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}

	if name, found := req.Form["name"]; found {
		fmt.Printf("Name is %v\n", name)
	}

	if age, found := req.Form["age"]; found {
		fmt.Printf("Age is %v\n", age)
	}
}
```

**启动前**注册 handler `http.HandleFunc` 对于每一个特定 URL，都需要一个对应的处理函数 `func(ResponseWriter, *Request)`

- `ResponseWriter` 响应输出
- `*Request` 请求
  - http.Request.URL.Path[1:] 表示去掉 URL 开头后的字符串子切片
  - e.g. /World → World
- http.HandlerFunc()

```go
http.HandleFunc("/", HelloServer)
http.HandleFunc("/submit", formHandler)

// start if err
http.ListenAndServe(":8080", http.HandlerFunc(HelloServer))
```

`fmt.Fprint` 和 `fmt.Fprintf` 都是用来写入 `http.ResponseWriter` 的不错的函数，实现了 io.Writer 接口。

```go
fmt.Fprintf(w, "<h1>%s<h1><div>%s</div>", name, age)
```

HTTP 使用 `http.ListenAndServeTLS()`

```go
http.Handle("/secure", http.HandlerFunc(secureHandler))

certFile := "/path/to/server.cert"
keyFile := "/path/to/server.key"

err := http.ListenAndServeTLS(":443", certFile, keyFile, nil)
```

类型重定义，将一个普通函数当作 HTTP 处理器的**适配器**。

```go
type HandlerFunc func(ResponseWriter, *Request)

// adaptor
http.Handle("/", http.HandlerFunc(HelloServer))

func HelloServer(w http.ResponseWriter, req *http.Request) {
    ...
}
```

结构体实现 `http.Handler`

```go
type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "struct Server")
}

func main() {
	var structServer Hello
	http.ListenAndServe(":8080", structServer)
}
```

