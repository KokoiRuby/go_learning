### Go Token

标记：关键字 keyword，标识符 identifier，常量 constant，字符串 string，符号 symbol

```bash
# 1. fmt
# 2. .
# 3. Println
# 4. (
# 5. "Hello, World!"
# 6. )
fmt.Println("Hello, World!")
```

### Line Separator

一行代表一个语句结束，无需 `;` 结尾

### Comment

```go
//  one-line comment

/*
	multi-line comment
*/
```

### Identifier

标识符用来**命名**变量、类型等程序实体。[Predeclared identifiers](https://go.dev/ref/spec#Predeclared_identifiers)

命名规范 Naming Convention:

- a-z/A-Z
- 数字不能开头
- Case-sensitive
- 不能有空格
- 非关键字

:construction_worker: Practice

- 每个 `.go` 属于一个包，包名和 `.go` 所在目录尽量保持一致。
- 驼峰命名
- 首字母大写则可以被其他包访问 = public

```bash
# valid
mahesh   kumar   abc   move_name   a_123
myname50   _temp   j   a23b9   retVal

# invalid
1ab  # starting with num
case # keyword
a+b  # operator not allowed
```

### [Keyword](https://go.dev/ref/spec#Keywords)

### operator

:warning: Go 不支持三元运算符，只支持二元。

#### Arithmetic

| 运算符 | 描述 | 实例 (A=10, B=20)  |
| :----: | :--: | :----------------: |
|   +    | 相加 | A + B 输出结果 30  |
|   -    | 相减 | A - B 输出结果 -10 |
|   *    | 相乘 | A * B 输出结果 200 |
|   /    | 相除 |  B / A 输出结果 2  |
|   %    | 求余 |  B % A 输出结果 0  |
|   ++   | 自增 |  A++ 输出结果 11   |
|   --   | 自减 |   A-- 输出结果 9   |

:construction_worker: Practice

- `/` = ceil 向下取整。
- `++`/`--` 只能单独使用，且只能写在变量后。

```go
var a int =5

a = i++ // nok
a = i-- // nok
a++     // ok
++a     // nok
```

#### Relational

| 运算符 | 描述                                                         | 实例              |
| :----- | :----------------------------------------------------------- | :---------------- |
| ==     | 检查两个值是否相等，如果相等返回 true  否则返回 false 。     | (A == B) 为 false |
| !=     | 检查两个值是否不相等，如果不相等返回 true  否则返回 false 。 | (A != B) 为 true  |
| >      | 检查左边值是否大于右边值，如果是返回 true  否则返回 false 。 | (A > B) 为 false  |
| <      | 检查左边值是否小于右边值，如果是返回 true  否则返回 false 。 | (A < B) 为 true   |
| >=     | 检查左边值是否大于等于右边值，如果是返回 true  否则返回 false 。 | (A >= B) 为 false |
| <=     | 检查左边值是否小于等于右边值，如果是返回 true  否则返回 false 。 | (A <= B) 为 true  |

#### Logical

摩根定律

- `!(a && b)` 等价于 `!a || !b`
- `!(a || b)` 等价于 `!a && !b`

| 运算符 | 描述                                                         |        实例        |
| :----- | :----------------------------------------------------------- | :----------------: |
| &&     | 逻辑 AND 运算符。 如果两边的操作数都是 true ，则条件 true ，否则为 false。 | (A && B) 为 false  |
| \|\|   | 逻辑 OR 运算符。 如果两边的操作数有一个 true ，则条件 true ，否则为 false。 | (A \|\| B) 为 true |
| !      | 逻辑 NOT 运算符。 如果条件为 true ，则逻辑 NOT 条件 false，否则为 true 。 | !(A && B) 为 true  |

#### Bit

| 运算符 | 描述                                                         | 实例                                   |
| :----- | :----------------------------------------------------------- | :------------------------------------- |
| &      | 按位与运算符"&"是双目运算符。 其功能是参与运算的两数各对应的二进位相与。 | (A & B) 结果为 12, 二进制为 0000 1100  |
| \|     | 按位或运算符"\|"是双目运算符。 其功能是参与运算的两数各对应的二进位相或 | (A \| B) 结果为 61, 二进制为 0011 1101 |
| ^      | 按位异或运算符"^"是双目运算符。 其功能是**参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。** | (A ^ B) 结果为 49, 二进制为 0011 0001  |
| <<     | 左移运算符"<<"是双目运算符。左移n位就是乘以2的n次方。 其功能把**"<<"左边的运算数的各二进位全部左移若干位，由"<<"右边的数指定移动的位数，高位丢弃，低位补0**。 | A << 2 结果为 240 ，二进制为 1111 0000 |
| >>     | 右移运算符">>"是双目运算符。右移n位就是除以2的n次方。 其功能是把**">>"左边的运算数的各二进位全部右移若干位**，">>**"右边的数指定移动的位数。** | A >> 2 结果为 15 ，二进制为 0000 1111  |