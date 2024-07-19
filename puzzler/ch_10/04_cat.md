**cat**

循环使用 os.Args 捕捉每一个参数中的文件名，打开并创建对应的 `bufio.NewReader` 然后逐行读取直到 EOF。

显示行数：通过定义 flag 变量控制是否显示行数，行数从 1 开始，每读取一行 ++。



**vs**

`os.File.Read([SIZE]byte)` 方法直接从文件中读取数据到用户提供的缓冲区。可根据提供的缓冲区大小，即**定长的字节切片**控制每次最多能读取的大小。适用于读取二进制或大文件分块读取。

`bufio.Reader('sep')` 在读取数据时使用内部缓冲区，减少了对底层文件系统或网络的直接读操作，通常性能更高。可一直读取到指定的分隔符并输出。适用于逐行读取。



**读写文件时使用缓冲有什么好处？**

1. 减少系统调用，可将多个小读写合并成一次大读写，降低开销。

2. 提高 IO 性能。

3. 减少磁盘访问次数。

4. 流处理。

5. 间接降低 CPU 使用率。

   

`fmt.Fprintf` 用于将格式化的字符串写入指定的 `io.Writer` 接口。`os.Stdout` 实现了 `io.Writer`。

记得 `Flush()` 会将缓冲区中的数据强制写入 `io.Writer`

```go
buf := bufio.NewWriter(os.Stdout)
fmt.Fprintf(buf, "%s\n", "hello world! - buffered")
buf.Flush()
```
