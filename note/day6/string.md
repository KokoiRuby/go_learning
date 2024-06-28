Common

| func                                   | Desc                             |
| -------------------------------------- | -------------------------------- |
| `len(str)`                             | 返回字符串长度 O(1)              |
| `range []rune(str)`                    | 遍历字符串转换成 rune/int32      |
| `strconv.Atoi(str)`                    | 数字字符串转整数                 |
| `strconv.Itoa(...)`                    | 整数转字符串                     |
| `[]byte(str)`                          | 字符串转字节                     |
| `string([]byte{...})`                  | 字节转字符串                     |
| `strconv.FormatInt(i int64, base int)` | 十进制数转2/8/16进制字符串       |
| `strings.Contains`                     | 判断是否包含子串                 |
| `strings.Count`                        | 统计字串出现次数                 |
| `str1 == str2`                         | 判断字符串是否相等               |
| `strings.EqualFold`                    | 判断字符串是否相等，不区分大小写 |
| `strings.Replace`                      | 替换字符串                       |
| `strings.ToUpper/ToLower`              | 大小写转换                       |
| `strings.Split`                        | 按 sep 进行拆分，返回一个切片    |
| `strings.Join`                         | 以 sep 进行组合切片成一个字符串  |
| `strings.Trim*`                        | 裁剪                             |
| `HasPrefix`<br />`HasSuffix`           | 是否有前后缀                     |

