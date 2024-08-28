使用模板 Template 去**合并数据结构与 html 中的数据**，对于构建 web app 非常拥有。

`{{.FieldName}}` 字段替换，其中 FieldName 是结构体中的一个字段 or map 中的 key。

```go
t := template.New("hello")
t, _ = t.Parse("hello {{.Name}}!")
t, _ = t.Parse("hello {{.}}!")

p := &Person{
	Name: "Alice",
	Age:  "18",
}

err := t.Execute(os.Stdout, p)
```

当模板应用在浏览器中时，要先用 html 过滤器去过滤输出的内容。

告诉 template 引擎在输出 FieldName 的值之前要通过 html 格式化它，进行特殊字符转义比如，防止数据破坏表单。

```go
// pipe-like
t, _ = t.Parse("hello {{.Name |html}}!")
```

使用 `Must` 验证模板语法是否正确。

```go
t := template.New("ok")
template.Must(t.Parse("/* a comment */ some static text: {{ .Name }}"))

t = template.New("nok")
// missing }
template.Must(t.Parse("some static text {{ .Name }"))
```

Commonly used

```go
var strTempl = template.Must(template.New("Name").Parse(strTemplateHTML))
```

### Pipeline

**管道**（Pipeline）是一种用于在模板中将多个函数或方法串联起来处理数据的语法结构。

它类似于 Unix 命令行中的管道概念，将一个函数的输出作为下一个函数的输入。

```go
{{ pipeline }}
```

使用 `if/else` 控制管道的输出

```go
tEmpty := template.New("template test")
tEmpty = template.Must(tEmpty.
	Parse("Empty pipeline if demo: {{if ``}} Will not print. {{end}}\n"))
tEmpty.Execute(os.Stdout, nil)

tWithValue := template.New("template test")
tWithValue = template.Must(tWithValue.
	Parse("Empty pipeline if demo: {{if `anything`}} Will not print. {{end}}\n"))
tWithValue.Execute(os.Stdout, nil)

tWithValue = template.New("template test")
tWithValue = template.Must(tWithValue.
	Parse("Empty pipeline if demo: {{if ``}} Will not print. {{else}} Print ELSE part. {{end}}\n"))
tWithValue.Execute(os.Stdout, nil)
```

`with` 语句将点的值设置为管道的值。如果管道是空的，就会跳过 `with` 到 `end` 之前的任何内容。

```go
// hello → .
"{{with `hello`}}{{.}}"
```

变量名前加一个 `$` 来为模板中的管道创建一个**局部变量**。

```go
"{{with $3 := `hello`}}{{$3}}{{end}}!\n"
```

`range` 在循环的集合中使用，管道的值必须是一个数组、切片或者 map。

如果管道的值的长度为零，点不会被影响并且 T0 会被执行

否则将点设置成拥有连续元素的数组、切片或者 map，T1 就会被执行。

```go
{{range pipeline}} T1 {{else}} T0 {{end}}
```

```go
`
<ul>
{{ range . }}
    <li>{{ .Name | ToUpper }} - {{ .Age }} years old</li>
{{ end }}
</ul>
`

type User struct {
	Name string
	Age  int
}

t, err := template.New("example").Funcs(template.FuncMap{
		"ToUpper": strings.ToUpper,
	}).Parse(tmpl)

users := []User{
		{Name: "Alice", Age: 23},
		{Name: "Bob", Age: 29},
		{Name: "Charlie", Age: 34},
	}
```

预定义模板函数 printf

```go
"{{with $x := `hello`}}{{printf `%s %s` $x `Mary`}} {{end}}!\n"
```

