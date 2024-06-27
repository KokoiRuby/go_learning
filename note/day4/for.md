for 循环有 3 种形式 where

- init： 一般为赋值表达式，给控制变量赋初值；（可选）
- **condition： 关系表达式或逻辑表达式，循环控制条件；**
- post： 一般为赋值表达式，给控制变量增量或减量。（可选）

```go
for init; condition; post { }
for condition { }
for { } // infinite loop
```

for range 支持 iteration

```go
for key, value := range oldMap {
    newMap[key] = vgoalue
}
```

:warning: Go 中是没有 while 和 do while 的语法，可通过 for 实现

```go
// while
// init
for{
	if bool_expr {
		break    
	}
	// do something
	// post
}
```

```go
// do while
// init
for{
	// do something
	// post
	if bool_expr{
		break
	}
}
```

