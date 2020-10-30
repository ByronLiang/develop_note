# TCP 粘包问题笔记

## 粘包背景

TCP是一个基于字节流的传输服务，"流"意味着TCP所传输的数据是没有边界的. 

这不同于UDP提供基于消息的传输服务，其传输的数据是有边界的. TCP的发送方无法保证对等方每次接收到的是一个完整的数据包

一个完整的应用层数据被分割成多次发送，导致接收对等方不是按完整数据包的方式来接收数据

### 无需考虑粘包的连接请求

1. 类似http的请求就不用考虑粘包的问题，因为服务端收到报文后, 就将缓冲区数据接收, 然后关闭连接;

2. 如果发送数据无结构，如文件传输，这样发送方只管发送，接收方只管接收存储就无需考虑粘包

### 粘包发送情景

双方建立连接，需要在连接后一段时间内发送不同结构数据(无法分辨消息与消息之间的边界在哪)

## 解决思路

1. 发送定长包。如果每个消息的大小都是一样的，那么在接收对等方只要累计接收数据，直到数据等于一个定长的数值就将它作为一个消息。

2. 包头加上包体长度。包头是定长的4个字节，说明了包体的长度。接收对等方先接收包体长度，依据包体长度来接收包体 【优先方案】

3. tcp每次发送数据，就与对方建立连接，然后双方发送完一段数据后，就关闭连接

### Design & Code

假设现在要发送[0x11, 0x22, 0x33]，约定协议头为[0xaa, 0xbb]，由于发送数据的长度是三个字节，所以经过客户端封装之后的数据就变成了[0xaa, 0xbb, 0x03, 0x11, 0x22, 0x33]

服务端收到数据后，先找[0xaa, 0xbb]的位置，然后根据他们的位置得到数据长度为3，于是再往后读三个字节就是真正的的数据部分了

```go
type PkgHeader struct {
	HeaderFlag	[2]byte
	DataLength	uint32
}

// 封装byte数据为二进制
buff := bytes.NewBuffer([]byte{})
binary.Write(buff, binary.BigEndian, []byte{0xaa, 0xbb}) //添加协议头
binary.Write(buff, binary.BigEndian, uint32(len(data))) //添加长度
binary.Write(buff, binary.BigEndian, data) //数据部分
allBytes := buff.Bytes()

// 解析二进制数据buffer
buff := bytes.NewBuffer(data)
binary.Read(buff, binary.BigEndian, &h.HeaderFlag) //读取0xaa 0xbb连个字节
binary.Read(buff, binary.BigEndian, &h.DataLength) //读取四个字节的长度
```

#### binary package

以Buffer为载体，将byte与二进制进行相互转换； `binary.Write()` 进行二进制数据的封装

