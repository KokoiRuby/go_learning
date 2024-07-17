## [Regexp](https://pkg.go.dev/regexp)

正则表达式 Regexp 是一种**描述文本模式 pattern** 的语言，使用特定的语法和规则来定义匹配字符串的规则。

通过正则表达式，可以**搜索、替换、验证和解析文本中的特定子串**。

### Syntax

1. **字符类**：

   - `[abc]`：匹配 'a'、'b' 或 'c' 中的任意一个字符。
   - `[^abc]`：匹配不在 'a'、'b' 和 'c' 中的任意字符。
   - `[a-z]`：匹配任意一个小写字母。
   - `[A-Z]`：匹配任意一个大写字母。
   - `[0-9]`：匹配任意一个数字字符。

2. **预定义字符类**:

   - `.`：匹配除换行符以外的任意单个字符。
   - `\d`：匹配任何一个数字字符，相当于 `[0-9]`。
   - `\D`：匹配任何一个非数字字符，相当于 `[^0-9]`。
   - `\w`：匹配任何一个字母、数字或下划线字符，相当于 `[a-zA-Z0-9_]`。
   - `\W`：匹配任何一个非字母、数字或下划线字符，相当于 `[^a-zA-Z0-9_]`。
   - `\s`：匹配任何一个空白字符（包括空格、制表符等）。
   - `\S`：匹配任何一个非空白字符。

3. **量词**:

   - `*`：匹配前面的字符零次或多次。

     - `a*`：匹配零个或多个 'a'。

   - `+`：匹配前面的字符一次或多次。

     - `a+`：匹配一个或多个 'a'。

   - `?`：匹配前面的字符零次或一次。

     - `a?`：匹配零个或一个 'a'。

   - `{n}`：匹配前面的字符恰好 n 次。

     - `a{3}`：匹配三个 'a'。

   - `{n,}`：匹配前面的字符至少 n 次。

     - `a{2,}`：匹配至少两个 'a'。

   - `{n,m}`：匹配前面的字符至少 n 次，至多 m 次。

     - `a{2,4}`：匹配二到四个 'a'。

4. **位置匹配**：

   - `^`：匹配字符串的开头。

   - `$`：匹配字符串的结尾。

   - `\b`：匹配一个单词边界。

     - `\bword\b`：匹配完整的单词 'word'。

   - `\B`：匹配非单词边界。

5. **分组和捕获**:
   - `()`：将括号内的表达式作为一个分组，匹配分组内容并捕获。
   
     - `(abc)`：匹配并捕获 'abc'。
   
   - `(?:...)`：非捕获组，只进行匹配但不捕获。
   
     - `(?:abc)`：匹配 'abc' 但不捕获。
   
   - `(?=...)`：正向前瞻，匹配前面的位置，后面的内容必须匹配。
   
     - `a(?=b)`：匹配 'a'，前提是后面跟着 'b'。
   
   - `(?!...)`：负向前瞻，匹配前面的位置，后面的内容必须不匹配。
   
     - `a(?!b)`：匹配 'a'，前提是后面不跟着 'b'。

```bash
# email
[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}

# phone#
\d{3}-\d{3}-\d{4}

# ip
\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b

# URL
https?://[^\s/$.?#].[^\s]*

# yyyy-mm-dd
\b\d{4}-\d{2}-\d{2}\b

# HH:MM:SS
\b\d{2}:\d{2}:\d{2}\b
```

### Common

[Regexp.FindString](https://pkg.go.dev/regexp#example-Regexp.FindString)：输入文本中查找**第一个匹配**正则表达式的字符串，并返回该匹配字符串。

[Regexp.FindAllString](https://pkg.go.dev/regexp#example-Regexp.FindAllString)：输入文本中查找**所有匹配**正则表达式的字符串，返回该匹配字符串。

Regexp.FindAllStringIndex：输入文本中查找**所有匹配**正则表达式的字符串，返回该匹配字符串的索引。

[Regexp.ReplaceAllString](https://pkg.go.dev/regexp#example-Regexp.ReplaceAllString)：替换输入文本中查找**所有匹配**正则表达式的字符串，返回替换后的文本。

[Regexp.FindStringSubmatch](https://pkg.go.dev/regexp#example-Regexp.FindStringSubmatch)：查找第一个匹配正则表达式的字符串及其子匹配项，返回一个包含匹配项的切片。

[Regexp.Expand](https://pkg.go.dev/regexp#example-Regexp.Expand)：使用模板字符串生成新的字节切片。

[Regexp.ExpandString](https://pkg.go.dev/regexp#example-Regexp.ExpandString)：使用模板字符串生成新的字节切片，处理字符串。
