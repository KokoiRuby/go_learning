Go 允许在条件判断语句中允许申明一个变量，该变量只能使用短变量声明 `:=`

```go
if bool_expr {
   // do something
}
```

```go
if bool_expr {
   // do something
} else {
  // do something
}
```

```go
if bool_expr {
   // do something
   if bool_expr {
      // do something
   }
}
```

