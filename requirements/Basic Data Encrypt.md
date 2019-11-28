# Data Encrypt

## API接口安全

- https协议流程

1. https的加密层(SSL/TLS)位于http层和tcp层；对于客户端(发送端)，把http的内容加密送到下层的TCP; 对于服务端(接收端)，负责将TCP送来的数据解密还原成http内容

2. 抓包工具在http层捕获的数据是明文；https不是对http报文进行加密，而是对业务数据进行加密，然后用http传输；

3. 抓包工具放置在两段https通道上是能拦截到明文数据

- 加密传输层

1. SSL（Secure Socket Layer，安全套接字层) SSL 协议位于 TCP/IP 协议与各种应用层协议之间，为数据通讯提供安全支持。
2. TLS（Transport Layer Security，传输层安全）

- 安全通信原则

1. 数据内容加密
2. 通讯双方身份校验
3. 数据内容完整性

### 接口数据传输加密方案：OpenSSL AES-128-CBC

```php
<!-- 保存生成的key 与 iv 密匙 用作客户端数据加密处理, 服务端解密处理 -->
$key = bin2hex(openssl_random_pseudo_bytes(8));
$iv = bin2hex(openssl_random_pseudo_bytes(8));
$encrypted = openssl_encrypt('hello world', 'aes-128-cbc', $key, 0, $iv);
$decrypted = openssl_decrypt($encrypted, 'aes-128-cbc', $key, 0 , $iv);
```

### 微信签名方式

将所有参数数据用 key => value 拼接起来，结尾接个 token，用 md5 加密后生成 sign 携带进来。
Api 端拿到参数后去校验 sign 用同样的方法去拼接，看看生成的 sign 和 前端传来的 sign 是否一致即可

## 常用加密算法

- 对称密码算法：是指加密和解密使用相同的密钥，典型的有DES、RC5、IDEA(分组加密)，RC4(序列加密)

- 非对称密码算法：又称为公钥加密算法，是指加密和解密使用不同的密钥（公开的公钥用于加密，私有的私钥用于解密）。比如A发送，B接收，A想确保消息只有B看到，需要B生成一对公私钥，并拿到B的公钥。于是A用这个公钥加密消息，B收到密文后用自己的与之匹配的私钥解密即可。反过来也可以用私钥加密公钥解密。也就是说对于给定的公钥有且只有与之匹配的私钥可以解密，对于给定的私钥，有且只有与之匹配的公钥可以解密。典型的算法有RSA，DSA，DH；

- 散列算法：散列变换是指把文件内容通过某种公开的算法，变成固定长度的值（散列值），这个过程可以使用密钥也可以不使用。这种散列变换是不可逆的，也就是说不能从散列值变成原文。因此，散列变换通常用于验证原文是否被篡改。典型的算法有：MD5，SHA，Base64，CRC等。
