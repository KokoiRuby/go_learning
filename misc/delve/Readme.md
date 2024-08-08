[Delve](https://github.com/go-delve/delve) is a **debugger** for the Golang.

### Installation

```bash
$ git clone https://github.com/go-delve/delve && cd delve
$ go install github.com/go-delve/delve/cmd/dlv
$ sudo ln -s $GOPATH/bin/dlv /usr/local/bin/dlv
```

### Quickstart

```bash
$ dlv debug main.go
```

