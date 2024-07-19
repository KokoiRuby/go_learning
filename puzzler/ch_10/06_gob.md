**gob**

一种用于编码和解码 Go 数据结构的二进制格式，主要用于 go 程序之间的数据传输，专用。

:smile: 较 XML/JSON 更小，二进制非文本，且少了 tag 和格式化信息。

:smile: 类型安全，可直接反序列化成原始类型，不需要转换。

Encode: `gob.NewEncoder(bytes.Buffer).Encode(&struct)` → []byte

Decode: `gob.NewEncoder(bytes.Buffer).decode(&struct_empty)` → Struct