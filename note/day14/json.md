JSON: JavaScript Object Notation JS 对象标记，通常用于**存储和交换数据**。

UC：Web/IM (Instant Messaging): Server → Serialize to JSON string → (transmission) → De-serialize from JSON string → Client parse.



**序列化 Marshal** = Go obj → JSON string。

```go
// struct
type S struct {}
s := S{}
jsonStr ,err := json.Marshal(&s)

// map
var m map[string]interface{}
m = make(map[string]interface{})
m["key"] = "value"
jsonStr ,err := json.Marshal(&m)

// slice
sl := []string{"Apple", "Banana", "Cherry"}
jsonStr3, err := json.Marshal(&sl)

// normal type
num := 123
jsonStr4, err := json.Marshal(num)
```

**反序列化 Unmarshal** = JSON string → Go obj。

- 对于 map/slice 不需要 make，make 被封装到了 Unmarshal 中

```go
json.Unmarshal([]byte(str), &s)
json.Unmarshal([]byte(str), &m)
json.Unmarshal([]byte(str), &sl)
```