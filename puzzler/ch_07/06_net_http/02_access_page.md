`http.Head` 发送 HEAD 请求，只返回响应 Header。

```go
res, err := http.Head(url)
```

`http.Get` 获取页面然后通过 `io.ReadAll` + `string` 读取

```go
res, err := http.Get(url)
data, err := io.ReadAll(res.Body)
string(data)
```

获取 twitter 用户状态，使用 `xml` 反序列化成一个结构体实例。

```go
type Status struct {
	Text string
}

type User struct {
	XMLName xml.Name
	Status
}

res, _ := http.Get("http://twitter.com/users/Googland.xml")
user := &User{
	xml.Name{
		"",
		"user",
	},
	Status{},
}

// get body in []byte
body, _ := io.ReadAll(res.Body)

xml.Unmarshal(body, user)
```

Const for HTTP status code.

```go
// net/http/status.go
const (
	...
	StatusOK                   = 200 // RFC 9110, 15.3.1
	StatusCreated              = 201 // RFC 9110, 15.3.2
    ...
}
```

使用 `w.header().Set ("Content-Type","../..")` 设置响应头信息

```go
type Response struct {
	Message string `json:"message"`
}

func headerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res := &Response{
		Message: "Hello",
	}

	// sercialize from struct to JSON
	json.NewEncoder(w).Encode(res)
}
```

