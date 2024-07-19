**encryption**

hash 包：实现了 adler32、crc32、crc64 和 fnv 校验；

crypto 包：实现了其它的 hash 算法，比如 md4、md5、sha1 等。以及完整地实现了 aes、blowfish、rc4、rsa、xtea 等加密算法。

md5 是一种常见的哈希函数，用于生成 128 位（16 字节）散列值，通常用于校验数据的完整性。

```go
hasher := md5.New()
io.Copy(hasher, file)               // cp file to hasher
hex.EncodeToString(hasher.Sum(nil)) // cal
```

