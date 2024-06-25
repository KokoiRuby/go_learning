#### Go Token

标记：关键字，标识符，常量，字符串，符号

```bash
# 1. fmt
# 2. .
# 3. Println
# 4. (
# 5. "Hello, World!"
# 6. )
fmt.Println("Hello, World!")
```

#### Line Separator

一行代表一个语句结束，无需 `;` 结尾

#### Comment

```go
//  one-line comment

/*
	multi-line comment
*/
```

#### Identifier

标识符用来**命名**变量、类型等程序实体。

一个标识符由若干个字母 (A-Z和a-z) 数字 (0-9)、下划线 _ 组成的序列。

第一个字符必须是字母或下划线而不能是数字。

```bash
# valid
mahesh   kumar   abc   move_name   a_123
myname50   _temp   j   a23b9   retVal

# invalid
1ab  # starting with num
case # keyword
a+b  # operator not allowed
```

#### [Keyword](https://go.dev/ref/spec#Keywords)

#### [Predeclared identifiers](https://go.dev/ref/spec#Predeclared_identifiers)