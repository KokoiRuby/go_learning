break 默认跳出最近的循环，多重循环则跳出 break 所在层的循环。

若想要跳到外层循环，后面可以后接标签 label，可以指定跳出到标签对应的循环中。

```go
break label
```

continue 跳过当前循环执行下一次循环。同样可使用 label，可用于跳转至外部循环。

```go
continue label
```

