# 加密协议应用事项

1. 对于非对称加密RSA, 针对bite为2048的公钥，加密消息长度不能大于256 - 11：`the length of the public modulus minus 11 bytes.`
   否则会发生`crypto/rsa: message too long for RSA public key size`加密消息的长度取决于公匙的长度
   bite为1024的公钥, 加密消息长度为128-11

2. 若需对过长消息进行加密处理，则需要分段截取消息长度进行加密；解析加密数据则按照顺序进行解析，并进行数据拼接

