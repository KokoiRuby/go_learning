## HTTP/2

### HTTP/1.1 性能问题

Website 变化：消息大小/页面资源/内容样式/实时性要求 ↑。

HTTP/1.1 性能问题：队头阻塞/头部开销（固定字段，重复，ASCII 编码效率低）/并发有限 (6 max Chrome)。

HTTP/1.1 优化：

- 图片二进制编码嵌入 HTML/CSS 减少网络请求次数。
- 资源分散到不同域名，提高并发数。

### 兼容

1. `http://` 明文 `https://` 加密
2. 基于 TCP，HTTP semantics 语义（不变） + syntax 语法（改变）

### 头部压缩

HPACK 算法：C/S 两端维护字典表 = 索引 Index → 字段；Huffman 编码。

**静态表**：高频 0~61。

**动态表**：62+ 编码解码时随时更新C/S 两端；后续只需发送 1 Byte Index。理论上 1 个连接内最终每个头部都会变成 1 Byte。

`http2_max_requests` 限制连接上请求数量避免动态表增大，超过连接关闭释放内存。

![image-20240105143006681](https://cdn.xiaolincoding.com//picgo/image-20240105143006681.png)

```bash
# HTTP/1.1
# 17 bytes
server: nghttpx\r\n   

# HTTP2
# 8 bytes
# char → Huffman
n  101010
g  100110
h  100111
t  001001
t  001001
p  101011
x  111001
```

### 帧

Headers + Data/Payload，二进制 instead of 明文，位操作效率更高。

- [Type](https://webconcepts.info/concepts/http2-frame-type/): 帧类型
- Flags: 帧标志控制位（优先级）
- **Stream Identifier: 标识该 Frame 属于哪个 Stream**

![enter image description here](https://i.sstatic.net/qwyOZ.png)

```bash
# status code: 200 

# HTTP/1.1
# 3 bytes
00110010 00110000 00110000 

# HTTP/2
# 1 byte only
10001000
```

### 并发

多 Stream 复用一个 TCP 连接；同一 HTTP 请求/响应在同一个 Stream 中；HTTP 请求/响应由多个 Frame 组成。

![image-20240105143239921](https://cdn.xiaolincoding.com//picgo/image-20240105143239921.png)

不同 Stream 的帧可乱序发送，即并发不同的 Stream，∵ Frame 携带 Stream ID，接收端最终可以拼凑出完整消息。

同一 Stream 内 Frame 有序。

**Stream ID (C奇S偶)**，唯一顺序递增，耗尽时 `GOWAY` 控制帧关闭 TCP。

### 主动推送

```bash
# HTTP/1.1
req html →
         ← res
req css  →
         ← css

# HTTP/2
req html →
         ← res + css 
```



req html →

​         ← res html + css 