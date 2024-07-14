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

:construction_worker: **Practice**

for range 遍历数组，∵ 数组属于值类型，所以传入的是副本，会进行浅拷贝。

```go
arr := [...]int{1, 2, 3, 4, 5, 6}
maxIndex := len(arr) - 1 // 5

for i, e := range arr {
	if i == maxIndex {
		arr[0] += e
	} else {
		arr[i+1] += e
	}
}

// [7 3 5 7 9 11]
fmt.Println(arr)
```

for range 遍历数组，∵ 数组属于引用类型，所以传入的是引用/指针，会进行深拷贝。

```go
sl := []int{1, 2, 3, 4, 5, 6}
maxIndex := len(sl) - 1 // 5
// 切片是引用对象，每一次迭代都会直接修改并反映到下一次迭代中
for i, e := range sl {
	if i == maxIndex {
		sl[0] += e
	} else {
		sl[i+1] += e
	}
}
// [22 3 6 10 15 21]
fmt.Println(sl)
```

