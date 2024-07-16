switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一，自上而下匹配，不需要 break。

匹配成功不会执行其他 case；若需跳过当前匹配 case 执行后面的 case，可以使用 fallthrough。

若都不匹配，执行 defaut。

**case 必须确保唯一性，否则编译不通过。但索引表达式，可以绕开。**

```go
// var1 could be any type
// val1 & val2 can be arbitrary value of var1 type
switch var1 {
    case val1:
        ...
    case val2, val3: // multi
        ...
    default:      
        ...
}

value := [...]int8{0, 1, 2, 3, 4, 5, 6}
switch expr { // nok due to duplicate case
	case 0, 1, 2:
		fmt.Println("0 or 1 or 2")
   	case 2, 3, 4:
		fmt.Println("2 or 3 or 4")
    
value := [...]int8{0, 1, 2, 3, 4, 5, 6}
switch expor { // ok
	case value[0], value[1], value[2]:
		fmt.Println("0 or 1 or 2")
   	case value[2], value[3], value[4]:
		fmt.Println("2 or 3 or 4")  
```

类型匹配判断 **interface{} 类型变量**的值。

:warning: 此时不支持 `fallthrough`。

**UC：在处理来自于外部类型未知的数据时，比如解析 JSON or XML，类型测试和转换会非常有用。**

```go
switch x.(type) {
    case type1:
       // do something      
    case type2:
       // do something  
    default:
       // do something 
}
```

:construction_worker: Practice

- 如果判断的具体数值不多，而且符合整数、浮点数、字符、字符串这几种类型，建议使用 switch，简洁高效。
- 如果是对**区间和结果为 bool 类型的判断，使用 if 更好**。
- 只有类型相同的值之间才有可能被允许进行判等操作。

```go
value := [...]int8{0, 1, 2, 3, 4, 5, 6}
switch 1 + 3 { // nok, int != int8
    case value[0], value[1]:
    	...
}
```

