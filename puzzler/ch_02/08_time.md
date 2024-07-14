### [time](https://golang.google.cn/pkg/time/)

`time` 包为我们提供了一个数据类型 `time.Time`（作为值使用）以及显示和测量时间和日期的功能函数。

可以使用 `time.Now()` 获取当前时间，或者使用 `t.*()`来获取时间的一部分。

`func (t Time) Format(layout string) string` 可以根据一个**格式化字符串**来将一个时间 t 转换为相应格式的字符串。

```go
import (
    "fmt"
    "time"
)

t := time.Now()
fmt.Println(t)
fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())

t = time.Now().UTC()
fmt.Println(t)

week = 60 * 60 * 24 * 7 * 1e9 // in nano sec
week_from_now := t.Add(week)
fmt.Println(week_from_now)

fmt.Println(t.Format(time.RFC822))
fmt.Println(t.Format(time.ANSIC))

fmt.Println(t.Format("21 Dec 2011 08:52")) // 21 Dec 2011 08:52
s := t.Format("20060102")
fmt.Println(t, "=>", s)
```

#### Sleep

`time.Sleep` 接受一个参数，表示要暂停的时间段。

时间常量：unix/unixnano timestamp (the number of s elapsed since Jan/1/1970 UTC) 只允许整数，不允许小数。

```go
const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)

time.Sleep(time.Second * 0.1)      // nok
time.Sleep(time.Millisecond * 100) // ok
```

#### After

#### Ticker
