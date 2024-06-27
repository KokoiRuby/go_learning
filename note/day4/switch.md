switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一，自上而下匹配，不需要 break。

匹配成功不会执行其他 case；若需跳过当前匹配 case 执行后面的 case，可以使用 fallthrough。

若都不匹配，执行 defaut。

```go
// var1 could be any time
// val1 & val2 can be arbitrary value of var1 type
switch var1 {
    case val1:
        ...
    case val2, val3: // multi
        ...
    default:      
        ...
}
```

类型匹配判断 **interface 类型的变量**的类型

```go
switch x.(type){
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