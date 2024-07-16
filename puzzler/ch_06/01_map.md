map 是键值对的集合。K as Index → V。K 不可重复，K 有序 since v1.12。

只有能通过 `==` 判断大小的数据类型才能作为 K；比如 slice/map/func 不可作 K。

:construction_worker: 线程不安全，并发请使用 `sync.Map`

对于结构体，必须要提供 `Key()` & `Hash()` 方法。

```go
type Person struct {
	Name string
	Age  int
}

func (p Person) Key() string {
	return p.Name
}

func (p Person) Hash() uint32 {
	h := fnv.New32a()
	_, err := h.Write([]byte(p.Key()))
	if err != nil {
		return 0
	}
	return h.Sum32()
}
```

map 是**引用类型**，传递的是引用（开销小），所以修改会直接影响。

:warning: 如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对。

:warning: map 声明是不会分配内存的，初始化需要 make，分配内存后才可以赋值与使用。

**:warning: 扩容原理：if load_factor >= 0.7，即 len/cap >= 0.7 时，扩 2 倍**

- **当 map 进行扩容时，会创建一个新的哈希表，并对所有 K 重新哈希，放到新的桶中** → 开销

```go
// 1
var myMap map[key_data_type]value_data_type
myMap := make(map[key_data_type]value_data_type[, len][, cap])
myMap["key"] = "value"

// 2
myMap := map[string]string{
	"key": "value",
}

// 3
myMap := make(map[string]string[, len])
myMap["key"] = "value"

// len only
len(myMap)
```

删除

```go
delete(m, "k")
m = make(map[string]string) // flush
```

判断是否包含 k

```go
if v, ok := m["k"]; ok {
    fmt.Println(v)
} else {
    fmt.Println("key not found")
}
```

迭代

```go
m["k1"] = "v1"
m["k2"] = "v2"
m["k3"] = "v3"
for k, v := range m {
	println(k, v)
}
```

map 切片，即切片元素是一个 map

```go
// make map slice
msl := make([]map[string]string, 1)
// chk if nil
if msl[0] == nil {
	// make map (in the slice)
	msl[0] = make(map[string]string, 3)
	msl[0]["k1"] = "v1"
	msl[0]["k2"] = "v2"
	msl[0]["k3"] = "v3"
}
fmt.Println(msl) // [map[k1:v1 k2:v2 k3:v3] map[]]

// append if cap is reached
msl = append(msl, newMap)
```

**空 map**：未通过 make 初始化；读返回零值，写直接 panic。

```go
var m map[string]int
value := m["key"]  // 0
m["key"] = 10      // panic: assignment to entry in nil map
```

**哪些适合做 K？**求哈希和判等**速度越快越好**。

- string 越短越好
- number 越窄越好 int8 | unint8 占 8 位 = 1 Byte
- struct 对它的所有字段值求哈希并进行合并

不建议你使用这些高级数据类型作为字典的键类型，不仅哈希判等慢，且值多变数。