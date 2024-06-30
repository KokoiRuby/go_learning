map 是键值对的集合。K as Index → V。K 不可重复，k 有序 since v1.12。

只有能通过 `==` 判断大小的数据类型才能作为 K；比如 slice/map/function 不可作 K。

map 是引用类型，传递的是引用，所以修改会直接影响。

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
myMap := make(map[string]string[, len][, cap])
myMap["key"] = "value"

// len & cap
len(myMap)
cap(myMap)
```

删除

```go
delete(m, "k")
m = make(map[string]string) // flush
```

判断是否包含 k

```go
v, ok := m["k"]
if ok {
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

