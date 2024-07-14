尽管 Go 支持 goto，但尽量不要使用。可以无条件地转移到过程中指定的行。

一般和 if 搭配使用。可用来实现条件转移，构成循环，跳出循环体等。

```go
label_back:
	// do something
	goto label_back

	goto label_front
	// do something
label_front: 
	// do something
```



