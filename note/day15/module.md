#### [Tutorial](https://go.dev/doc/tutorial/create-module)

```bash
example.com/
 |-- greetings/
 |-- hello/
```

The `go mod init` command creates a `go.mod` file to **track your code's dep**.

```bash
# start module
$ go mod init example.com/greetings
$ go mod init example.com/hello
```

```bash
# example.go/greetings/go.mod
module example.go/greetings

go 1.20

# example.go/hello/go.mod
module example.go/hello

go 1.20
```

```go
// example.go/hello/hello.go
package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
```

```bash
# re-loc dep to local dir
$ go mod edit -replace example.com/greetings=../greetings
```

`go mod tidy` command to synchronize dep.

```bash
$ go mod tidy
```

