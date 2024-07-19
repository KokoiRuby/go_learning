**cmd args:**

1. os 包中有一个 string 类型的**切片变量** `os.Args`，可读取命令行输入的参数，其中
   - `os.Args[0]` 是程序名。
   - `os.Args[1+]` 是传入参数。
2. `flag` 包用于解析命令行参数。通过 `flag.Type(name, value, usage)` 定义，其中
   - `name` 是 flag 名，`value` 是默认值，`usage` 是描述。
   - 命令行 `-name val`。若未指定 flag 则使用默认值。
   - 使用 `-h` 可以打印所有的 `usage`。
   - flag 变量通过 `*` 操作符获取值。