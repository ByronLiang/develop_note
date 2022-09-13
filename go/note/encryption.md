# 加密算法补码处理

## 背景

1. 常用加密算法如 AES 等主要对字符串进行加密，而解码后，补码会遗留到明文处, 但不会显示处文本内容(补码为0字节数据, 非ASCII码内容, 如`\x00\x00`)

2. 若使用加密算法对数值浮点字符进行加密与解码，需要注意浮点精度与解码后因字符含补码，解析浮点数据结构异常

### 关联案例

[AES CBC ciphertext gets padded with 16 0x00 bytes](https://stackoverflow.com/questions/72499124/golang-go-aes-cbc-ciphertext-gets-padded-with-16-0x00-bytes-for-some-reason)

## 知识与措施

1. 使用 `strconv.Quote()` 能将非ASCII码内容显示出来, 可视化原理: `\x`拼接高四位(`字节>>4`) 与 低四位(`字节&0xF`)转换`[0-f]`

2. 当 `strconv.ParseFloat()` 将非浮点字符进行解析，会报`syntaxError`类型错误

3. 对解码后字符串进行移除补码: `newBytes := bytes.TrimRight(解码字节数组, "\x00")`
