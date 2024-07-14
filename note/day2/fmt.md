### vs

| 函数       | 主要用途                                         | 输出行为                               | 返回值               |
| ---------- | ------------------------------------------------ | -------------------------------------- | -------------------- |
| `Println`  | 打印一系列值，自动添加空格和换行符               | 每个值之间有空格，末尾有换行符         | 无返回值             |
| `Printf`   | 格式化输出，根据格式化动词指定输出格式           | 按照格式化动词指定的格式输出，无换行符 | 无返回值             |
| `Sprintln` | 打印一系列值，自动添加空格和换行符，但返回字符串 | 每个值之间有空格，末尾有换行符         | 返回格式化后的字符串 |

```go
name := "Alice"
age := 30
// Name: Alice Age: 30
fmt.Println("Name:", name, "Age:", age)

// Name: Alice Age: 30
fmt.Printf("Name: %s Age: %d\n", name, age)

result := fmt.Sprintln("Name:", name, "Age:", age)
// Name: Alice Age: 30
fmt.Print(result)
```

