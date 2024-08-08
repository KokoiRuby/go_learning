## Background

**对于一个业务实体比如服务器 Server，我们通常需要对其进行配置，并非所有配置都是必须的。**

```go
type Server struct {
    Addr     string
    Port     int
    Protocol string
    Timeout  time.Duration
    MaxConns int
    TLS      *tls.Config
}
```

首先最容易想到的就是通过**构造器**，但由于 Go **不支持重载**，对于不同的配置需要创建多个构造器，简直是个灾难...

```go
NewDefaultServer(addr string, port int) (*Server, error)
NewTLSServer(addr string, port int, tls *tls.Config) (*Server, error)
NewServerWithTimeout(addr string, port int, timeout time.Duration)
```

## Solution

### Config

将配置和业务实体进行组合，只需要一个构造器即可。

注意：配置不一定是必须的，在构造器中对配置进行**判 nil 处理**。

```go
type Config struct {
    Protocol string
    Timeout  time.Duration
    Maxconns int
    TLS      *tls.Config
}

type Server struct {
    Addr string
    Port int
    Conf *Config   // composite
}

NewServer(addr string, port int, conf *Config) (*Server, error)
```

### Option

类 Builder 模式：通过 `With*` 链式调用构造对象；**Options 模式更加简洁，不需要引用额外的 Builder 对象。扩展性 & 可读性 ↑**

1. 类型重定义一个函数类型。

```go
type Option func(*Server)
```

2. 创建一系列配置函数：传入一个参数，返回上述重定义的函数类型的匿名函数，通常是 setter。

```go
func Protocol(p string) Option {
    return func(s *Server) {
        s.Protocol = p
    }
}
func Timeout(timeout time.Duration) Option {
    return func(s *Server) {
        s.Timeout = timeout
    }
}
func MaxConns(maxconns int) Option {
    return func(s *Server) {
        s.MaxConns = maxconns
    }
}
func TLS(tls *tls.Config) Option {
    return func(s *Server) {
        s.TLS = tls
    }
}
```

3. 定义构造器：接受一个 Option 类型的切片，遍历逐一进行配置。

```go
func NewServer(addr string, port int, options ...Option) (*Server, error) {
	srv := Server{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		Timeout:  30 * time.Second,
		MaxConns: 1000,
		TLS:      nil,
	}
	for _, option := range options {
		option(&srv)
	}
	//...
	return &srv, nil
}
```

```go
s1, _ := NewServer("localhost", 1024)
s2, _ := NewServer("localhost", 2048, Protocol("udp"))
s3, _ := NewServer("0.0.0.0", 8080, Timeout(300*time.Second), MaxConns(1000))
```

