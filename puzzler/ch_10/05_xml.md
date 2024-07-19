`encoding/xml` 

**Encode(to)** (go struct → xml doc) vs. **Decode(from)** (xml doc → go struct)

需要定义结构体 & 字段 tag → `xml:"tag_name_in_xml"`

编码时指定 Indent 缩进 `encoder.Indent("", "  ")` 一般无前缀，缩进未 `\t`