`fmt.Scanln` 函数用于从 stdin 中读取输入，并根据提供的格式将输入存储到指定的变量中，直到换行符。

```go
var num int
fmt.Print("Enter a number: ")
_, err := fmt.Scanln(&num)
```

`fmt.Sscanf` 函数用于从**字符串**中读取输入，并根据提供的格式将输入存储到指定的变量中。

```go
str := "42 John"

var num int
var name string

n, err := fmt.Sscanf(str, "%d %s", &num, &name)
```

`bufio.Reader` 提供缓冲读取数据，可以提高效率，对大量数据友好。

```go
reader := bufio.NewReader(os.Stdin)
for {
    line, err := reader.ReadString('\n') // win \r\n
    fmt.Println("Input:", line)
    if line == "\n" {                    // win \r\n
			break
	}
}
```

