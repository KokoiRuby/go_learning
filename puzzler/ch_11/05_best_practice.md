### DOs

（1）尽可能的使用 `:=` 去初始化声明一个变量（在函数内部）；

（2）尽可能的使用**字符代替字符串**；

（3）尽可能的使用**切片代替数组**；

（4）尽可能的**使用数组和切片代替映射**；

（5）如果只想获取切片中某项值，不需要值的索引，尽可能的使用 `for range` 去遍历切片，这比必须查询切片中的每个元素要快一些；

（6）当数组元素是稀疏的（例如有很多 `0` 值或者空值 `nil`），使用映射会降低内存消耗；

（7）初始化映射时指定其容量；

（8）当定义一个方法时，使用指针类型作为方法的接受者；

（9）在代码中使用常量或者标志提取常量的值；

（10）尽可能在需要分配大量内存时使用缓存；

（11）使用缓存模板。

### string

1. 修改字符串字符，转换成字节切片，再转回来。

```go
str:="hello"
c:=[]byte(str)
c[0]='c'
s2:= string(c) // s2 == "cello"
```

2. substring

```go
substr := str[start:end]
```

3. iter

```go
for i:=0; i < len(str); i++ {
	_ = str[i]
}

for i, v := range str {
…
}
```

4. len

```go
// # of bytes
len(str)

// # of rune
utf8.RuneCountInString(str)

// to int slice
len([]int(str))
```

5. concat

```go
var buf bytes.Buffer
buf.WriteString("Hello, ")

// str slice
strs := []string{"Hello", "World!"}
result := strings.Join(strs, ", ")
```

### array & slice

1. create

```go
arr := new([len]type)
sl := make([]type, len)
```

2. init

```go
arr := [...]type{i1, i2, i3, i4, i5}
var sl []type = arr[start:end]
```

3. cut tail

```go
sl = sl[:len(sl)-1]
sl = arr[:len(arr)-1]
```

4. iter

```go
for i:=0; i < len(arr); i++ {
	_ = arr[i]
}
for i, v := range arr {
	_
}
```

5. find elem in matrix

```go
found := false
Found: for row := range arr2Dim {
    for column := range arr2Dim[row] {
        if arr2Dim[row][column] == V{
            found = true
            break Found
        }
    }
}
```

### map

1. create & init

```go
m := make(map[keytype]valuetype)
m := map[keytype]valuetype{k1: v1, k2: v2}
```

2. iter

```go
for k, v := range m {
    _
}
```

3. if had key

```go
if v, ok := m[k]; ok {
    _
}
```

4. del

```go
delete(m, k)
```

### struct

1. create

```go
type struct1 struct {
    field1 type1
    field2 type2
    …
}
ms := new(struct1)
```

2. init

```go
ms := &struct1{10, 15.5, "Chris"}
```

3. factory

```go
func Newstruct1(n int, f float32, name string) *struct1 {
    return &struct1{n, f, name} 
}
```

### interface

1. type assertion

```go
if v, ok := v.(Type); ok {
    _
}
```

2. type switch

```go
func classifier(items ...interface{}) {
    for i, x := range items {
        switch x.(type) {
        case bool:
            fmt.Printf("param #%d is a bool\n", i)
        case float64:
            fmt.Printf("param #%d is a float64\n", i)
        case int, int64:
            fmt.Printf("param #%d is an int\n", i)
        case nil:
            fmt.Printf("param #%d is nil\n", i)
        case string:
            fmt.Printf("param #%d is a string\n", i)
        default:
            fmt.Printf("param #%d’s type is unknown\n", i)
        }
    }
}
```

### func

内建函数封装 panic recovery

```go
func protect(g func()) {
    defer func() {
        log.Println("done")
        if x := recover(); x != nil {
            log.Printf("run time panic: %v", x)
        }
    }()
    g()
}
```

### file

1. read

```go
file, err := os.Open("/path/to/file")
defer func(f *os.File) {
    if err := f.Close(); err != nil {
        log.Printf("Error closing file: %v", err)
    }
}(file)

iReader := bufio.NewReader(file)

for {
    str, err := iReader.ReadString('\n')
    if err == io.EOF {
      return
    }
    fmt.Printf("The input was: %s", str)
}
```

2. read/write by byte slice

```go
func cat(f *file.File) {
  const NBUF = 512
  var buf [NBUF]byte
  for {
    switch nr, er := f.Read(buf[:]); true {
    case nr < 0:
    	fmt.Fprintf(os.Stderr, "cat: error reading from %s: %s\n",
                    f.String(), er.String())
  		os.Exit(1)
    case nr == 0: // EOF
    	return
    case nr > 0:
    	if nw, ew := file.Stdout.Write(buf[0:nr]); nw != nr {
        	fmt.Fprintf(os.Stderr, "cat: error writing from %s: %s\n", 
                        f.String(), ew.String())
      }
    }
  }
}
```

### goroutine

1. 带缓冲通道能够极大提升吞吐。

```go
ch := make(chan type, buf)
```

2. 限制一个通道的数据数量并将它们封装成一个数组，在接收端解压，可提升性能。

```go
ch := make(chan []int, 3)
for i := 0; i < len(data); i += 3 {
    end := i + 3
    if end > len(data) {
        end = len(data)
    }
    chunk := data[i:end]

    ch <- chunk
}

// recv
for chunk := range ch {
    fmt.Println("Received chunk:", chunk)
    // TODO
}
```

3. iter

```go
for v := range ch {
    _
}
```

4. check if closed

```go
// auto
for v := range ch {
    _
}

// alter
for {
    if v, open := <-ch; !open {
        break
    }
    //
}
```

5. main goroutine wait until goroutines

```go
ch := make(chan struct{})
go func() {
    ch <- sturct{}
}()
// wait for goroutine to finish
<-ch
```

6. factory

```go
func pump() chan int {
    ch := make(chan int)
    go func() {
        for i := 0; ; i++ {
            ch <- i
        }
    }()
    return ch
}
```

7. 通道迭代器：用于在并发环境中安全地遍历通道的数据，确保收发同步无竞争。基于 `sync.WaitGroup` 

```go
var wg sync.WaitGroup

wg.Add(1)
go sendData(ch, &wg)        // defer wg.Done()

wg.Add(1)
go channelIterator(ch, &wg) // defer wg.Done()

wg.Wait()
```

8. 限制并发请求数量。基于 `sync.WaitGroup` 

```go
func processRequest(ch chan bool, wg *sync.WaitGroup, id int) {
	defer wg.Done()

	// consume, block when ch is empty
	<-ch

	// sim proc
	fmt.Printf("Processing request %d...\n", id)
	time.Sleep(2 * time.Second)

	// to keep maxConcurrent, blocked when ch is full
	ch <- true
	fmt.Printf("Request %d processed.\n", id)
}

maxConcurrent := 3
ch := make(chan bool, maxConcurrent)
var wg sync.WaitGroup

for i := 1; i <= 10; i++ {
	wg.Add(1)

    go processRequest(ch, &wg, i)

	// block at the beginning, stuff 3 into
    if i <= maxConcurrent {
		ch <- true
	}
}

wg.Wait()
```

9. 多核并行，基于信号量模式 semaphore。

```bash
$ GOMAXPROCS=NCPU
```

```go
func DoAll() {
    sem := make(chan int, NCPU)
    // start NCPU goroutine to proc
    for i := 0; i < NCPU; i++ {
        go DoPart(sem)
    }
    // pop if goroutine is empty
    for i := 0; i < NCPU; i++ {
        <-sem // block if empty
    }
}

func DoPart(sem chan int) {

    // proc
    sem <- 1 // block if full
}

func main() {
    runtime.GOMAXPROCS = NCPU
    DoAll()
}
```

10. 中止 goroutine (TODO)

```go
runtime.Goexit()
```

11. timeout

```go
timeout := make(chan bool, 1)
go func() {
    time.Sleep(time.Seconds * 1)
    timeout <- true
}()

select {
case <- ch:
    //
case <-timeout:
    //
}
```

12. lock

```go
func Worker(in, out chan *Task) {
    for {
        t := <-in
        process(t)
        out <- t
    }
}
```

14. 同步调用运行时间过长时将之丢弃 (TODO)

```go

```

15. timer (TODO)

```go
```

16. server backed (TODO)

```go
```

### web

template

```go
var strTempl = template.Must(template.New("TName").Parse(strTemplateHTML))
```

cache

### Misc

exit

```go
if err != nil {
   fmt.Printf("Program stopping with error %v", err)
   os.Exit(1)
}

if err != nil { 
    panic("ERROR occurred: " + err.Error())
}
```

