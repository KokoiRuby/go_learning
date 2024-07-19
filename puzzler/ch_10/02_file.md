文件：**数据源**的一种，主要用于**保存数据**；形式可以是文本/图片/音频视频。

流 Stream：描述数据源（文件） & 程序（内存）之间的路径。**方向相对于程序/内存而言。**

- InputStream 输入流：程序（内存）← 数据源（文件）
- OutputStream 输出流：程序（内存）→ 数据源（文件）



![image-20240701132339953](./02_file.assets/image-20240701132339953.png)



```go
// read-only
file, err := os.Open("/path/to/file")
if err != nil {
	fmt.Println("Error opening file:", err)
	return
}
// file is a ptr = fd
fmt.Printf("file= %v", file)
defer file.Close()
```

### **Open**

`os.OpenFile` 允许指定打开文件的模式

- `os.O_RDONLY`：只读模式
- `os.O_WRONLY`：只写模式
- `os.O_RDWR`：读写模式
- `os.O_APPEND`：追加模式
- `os.O_CREATE`：如果文件不存在，则创建
- `os.O_TRUNC`：如果文件已存在，打开时清空文件
- `os.O_EXCL`：与 `os.O_CREATE` 一起使用，文件必须不存在

```go
file, err := os.OpenFile("./sample.txt", os.O_WRONLY|os.O_CREATE, 0666)
if err != nil {
	fmt.Println(err)
}
defer file.Close()
```

### **Read**

`file.Read` 一次性读取数据存入字节切片 as Buffer 返回读取到的字节数，可 make 自定义大小。

```go
buf := make([]byte, 1024)

for {
	n, err := file.Read(buf)
	if err != nil {
		if err == os.ErrClosed {
			fmt.Println("File closed.")
			break
		} else if err == io.EOF {
			if n > 0 {
				fmt.Print(string(buf[:n]))
			}
			fmt.Println("Reached end of file.")
			break
		} else {
			fmt.Println("Error reading file:", err)
			break
		}
	}
	fmt.Print(string(buf[:n]))
}
```

`os.ReadFile` 全部读取到内存中；若文件太大，可能会导致内存溢出 :warning:

```go
byteString, err := ioutil.ReadFile("/path/to/file")
fmt.Println(string(byteString))
```

`bufio.NewReader` 提供了高效的**缓冲**读取功能，可以**逐行读取**文件内容。:smile:

```go
reader := bufio.NewReader(file)

for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println(line)
				break
			}
			fmt.Println(err)
			break
		}
		fmt.Println(line)
	}
```

`encoding/csv` 读取以 `,` 分隔的 `.csv` 文件。

```go
import (
	"encoding/csv"
)
file, err := os.Open("/path/to/.csv")
reader := csv.NewReader(file)
records, err := reader.ReadAll()
```

`compress/[bzip2|flate|gzip|lzw|zlib]` 提供了读取压缩文件的功能。

```go
filePath := "compressed_file.ext"
file, err := os.Open(filePath)
var reader io.Reader
	switch {
	case isBzip2(filePath):
		reader = bzip2.NewReader(file)
    }
	...
}

buffer := make([]byte, 1024)
n, err := reader.Read(buffer)
content := string(buffer[:n])

func isBzip2(filePath string) bool {
	return filepath.Ext(filePath) == ".bz2"
}
```

:bookmark_tabs: **对于超大文件该如何处理？**

1. **流处理**：在读取数据时使用内部缓冲区，减少了对底层文件系统或网络的直接读操作，通常性能更高。可一直读取到指定的**分隔符**并输出。适用于**逐行读取**。

```go
f, err := os.Open(filePath)
buf := bufio.NewReader(f)
line, err := buf.ReadLine("\n")
```

2. **分片处理**：从文件中读取数据到用户提供的缓冲区。可根据提供的缓冲区大小，即**定长的字节切片**控制每次最多能读取的大小。适用于**无换行符**的**读取二进制或大文件分块**读取。

```go
f, err := os.Open(fileName)
s := make([]byte, 4096)
switch nr, err := f.Read(s[:]);
case nr < 0:
	// err
case nr == 0:
	// EOF
case > 0:
	// TODO
```

**读写文件时使用缓冲有什么好处？**

1. 减少系统调用，可将多个小读写合并成一次大读写，降低开销。

2. 提高 IO 性能。

3. 减少磁盘访问次数。

4. 流处理。

5. 间接降低 CPU 使用率。

### **Write**

`buffio.NewWriter` 提供了一个高效的方式来进行**缓冲**写入操作，通过减少实际的 I/O 操作次数来提高写入性能，**最后须通过 flush 落盘。**

```go
file, err := os.OpenFile("/path/to/write", os.O_WRONLY|os.O_CREATE, 0666)
file, err := os.OpenFile("/path/to/write", os.O_RDWR|os.O_APPEND, 0666)

writer := bufio.NewWriter(file)
writer.WriteString(str)
writer.Flush()
```

### **Stat**

`os.Stat()` 

```go
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, nil
}
```

### **Copy**

1. 打开源文件
2. 创建目标文件
3. 创建 reader & writer
4. 循环读取源文件并写入目标文件
5. writer flush
6. 关闭源文件 & 目标文件

```go
io.Copy(writer, reader)
```

**统计**

**逐行遍历**统计该行有多少个 英文、数字、空格和其他字符，并保存到一个结构体。

```go
type CharCount struct {
	az    int
	num   int
	space int
	other int
}

line, err := reader.ReadString('\n')

for _, v := range line {
    swtich {
    	// TODO
    }
}
```

### Close

defer 确保返回是一定会执行，匿名函数用于在处理在关闭处理时可能发生的错误。

```go
// nok
defer file.Close()

// better
defer func() {
	if err := file.Close(); err != nil {
		fmt.Println("Error closing file:", err)
	}
}()
```

