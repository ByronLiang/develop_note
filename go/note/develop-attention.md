# 日常开发总结

## JSON 时如何区分空字段和未设置字段

解决方法: 将结构体的字段成员设置为指针类型

- 原始 JSON 中不存在该字段，则结构体里该成员字段将为空 (nil).

- 如果该字段确实存在并且为空字符串，则指针不为空，并且该字段包含空值("")

### 关注点

1. 非指针数据类型具备固有的空安全性。在字符串或整型永远不能为空。他们始终具备默认值

2. 指针数据类型在未手动设置的情况下默认为空, 读取空指针成员将发送异常; 应先判断是否空指针, 再访问成员字段数据

## 接口值 interface

1. 接口值，由两个部分组成，一个具体的类型(动态类型)和类型的值(动态值)

2. 一个包含nil指针(具备类型，但类型的值为nil)的接口不是nil接口(类型与类型的值都为nil)

## 类型比较

### 基础类型

比较的两个变量类型必须相等。由于没有隐式类型转换，比较的两个变量必须类型完全一样，类型别名也不行。如果要比较，先做类型转换再比较。

1. 类型完全不一样的，不能比较
2. 类型再定义关系，不能比较，可以强转比较
3. 类型别名关系，可以比较

## 相关总结

1. 复合类型，只有每个元素(成员)可比较，而且类型和值都相等时，两个复合元素才相等
2. slice，map不可比较，但是可以用reflect或者cmp包来比较
3. func作为一个类型，也不能比较;
4. interface{}类型变量比较: 只有`动态类型`和`动态值`都相同时，两个接口变量才相同
5. 引用类型的比较是看指向的是不是同一个变量
6. 类型再定义(type A string)不可比较，是两种不同的类型
7. 类型别名(type A = string)可比较，是同一种类型

## 汉字拼音排序处理

主要依赖`golang.org/x/text`库，涉及包: `simplifiedchinese` 与 `transform`

对汉字进行转码处理 UTF8转换为GBK 再根据拼音字符字母排序

```go
//import (
//    "golang.org/x/text/encoding/simplifiedchinese"
//    "golang.org/x/text/transform"
//)

type SortByPinyin []string

func (s SortByPinyin) Len() int {
	return len(s)
}

func (s SortByPinyin) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortByPinyin) Less(i, j int) bool {
	a, _ := UTF82GBK(s[i])
	b, _ := UTF82GBK(s[j])
	bLen := len(b)
	for idx, chr := range a {
		if idx > bLen-1 {
			return false
		}
		if chr != b[idx] {
			return chr < b[idx]
		}
	}
	return true
}

//UTF82GBK : transform UTF8 rune into GBK byte array
func UTF82GBK(src string) ([]byte, error) {
	GB18030 := simplifiedchinese.All[0]
	return ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), GB18030.NewEncoder()))
}

//GBK2UTF8 : transform  GBK byte array into UTF8 string
func GBK2UTF8(src []byte) (string, error) {
	GB18030 := simplifiedchinese.All[0]
	bytes, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader(src), GB18030.NewDecoder()))
	return string(bytes), err
}

// sample
user := []string{"张三", "李四"}
sort.Sort(SortByPinyin(user))
```

### 英文字符标准化处理

针对多音节英文字符，如`Pécs` 需要使用 rune 数据结构存储(32bit), 若标准化输出: `Pecs`，只需4bit存储

```go

import (
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var normalizer = transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)

func normalize(str string) (string, error) {
	s, _, err := transform.String(normalizer, str)
	if err != nil {
		return "", err
	}
	return strings.ToLower(s), err
}
```

参考：[字符标准化](https://go.dev/blog/normalization)

## 切片应用优化

- 对于可预期长度的切片初始化, 使用`temp := make([]int, 0, n)` 配置切片的容量cap长度，减少append操作, 扩容导致的旧数据拷贝

- 针对切片类型的数据作为函数的形参传递，需要以指针类型进行传值；`TempFun(n *[]int)` 函数内即便发生扩容, 切片数据仍保持有效

## nil值理解

- nil是一种变量；适用于指针，函数，interface，map，slice，channel这6种类型

- 针对变量定义 `var set map[int]string` 分配变量指定大小的内存，且都为置0分配, 并且确定一个变量名称

- 凡是变量定义, 它与nil进行值判断都是为true; nil值判断, 主要判断当前变量里的内存是否为0分配

- 针对切片slice, 变量定义与`make([]string, 0)`具有相同效果[定义`struct slice` 核心结构已分配内存]；而其余类型而是有差异的

## defer值返回


```go
// 非命名值返回
func deferHandle(i *int) int {
	defer func() {
		*i++
	}()
	return *i
}

// 命名值返回
func deferHandleWithName(i *int) (m int) {
    m = *i
    defer func() {
    	// 影响最终返回的m值
        m++
    }()
    return m + 10
}
```

- `defer语句`在`return语句`之后执行

- 非命名返回值：取决于return时候的值

- 命名返回：如同函数内的全局变量, 先执行完return的值逻辑，再执行defer对值逻辑; 执行函数体内对命名返回值的任何修改

## 针对Post请求体涉及数组解析 字段异构解析

msgBody的数组里两个元素组成；包含text与location 两个文本与坐标的异构体

```json
{
  "msgBody": [
    {
      "type": "text",
      "content": {
        "text": "something"
      }
    }, {
      "type": "location",
      "content": {
        "desc": "someone",
        "latitude": 129.340656774469956,
        "longitude": 116.77497920478824
      }
    }
  ]
}
```

针对http里的`req *http.Request` 对请求体`req.Body`转换成struct: 对于字段是数组类型，或者异构体，解析的成员类型需要是`interface{}`，否则会解析为空

异构体：当type字段是不同值时，content字段呈不同数据结构; 

```go
package main

// 针对整个数组数据进行解析
type Body struct {
	MsgBody interface{} `json:"msgBody"`
}

// []MsgBodyEle 切片类型
type MsgBody struct {
	Content []*MsgBodyEle `json:"msgBody"`
}

// 针对数组里每个元素的异构体进行解析
type MsgBodyEle struct {
	Type string `json:"type"`
	Content interface{} `json:"content"`
}
```

### 兼容与通用性好的处理方式

针对同一字段动态数据结构，字段数据结构使用 `json.RawMessage` 在序列化成功后，再对此字段依据相关标识，反序列化成相应对象

```go
type MsgBodyEle struct {
	Type string `json:"type"`
	Content json.RawMessage `json:"content"`
}

type TextContent struct {
	Text string `json:"text"`
}

var demo MsgBodyEle
if demo.Type == "text" {
	var contentText TextContent
	// RawMessage 可转换成 json 字节, 或直接进行反序列化
	// // 指定对Content 反序列成 文本对象
	json.Unmarshal(demo.Content, &contentText)
}
```

## for-select-default 模式引发 CPU 空转原理 (高CPU占用)

```go
for {
	select
	case xx:
	default:
	// continue 引发 for 空转
}
```
若 default 没有阻塞，则会导致不断进行 for 循环，因无法阻塞，导致长期占用 CPU

`for-select` 各个分支都处于阻塞状态，程序会进行调度，使资源充分发挥作用

只要default分支有一定任务运行或具有阻塞特性，则不会引发长期占用 CPU 的情况。但需要确保执行的任务，不能是for无限循环，否则也会引发 高CPU 负载的情景出现

## fmt.Sscanf

对字符串按照格式提取内容, 并存储到指定变量里

如: 
```go
var (
	s string
	n int
)
// 报错: unexpected EOF
fmt.Sscanf("demo-125.txt", "%s.txt", &s)
fmt.Sscanf("demo-125.txt", "%s-%d.txt", &s, &n) // 由于没有空格进行分隔，第一个变量获取全部字符内容
// 成功提取: 参数125
fmt.Sscanf("demo-125.txt", "demo-%d.txt", &n)
// fmt.Sscanf("demo 125.txt", "%s %d.txt", &s, &n) 能成功提取 demo 字符内容与 125 数值
```

### 出现 unexpected EOF 原因

若针对字符数据结构，进行解析，被提取内容默认需要以空格或制表符(\t)等方式与原文内容分离，才能正常提取

主要机制: 内部对字符内容的分割原理，需要针对空格进行分割, 而机制(`fmt.Scanner`) 是个接口，只需实现接口，使其支持以不同方式进行兼容分割

### 解决

对字符内容的分割原理主要依赖机制(`fmt.Scanner`) 它是个接口，只需实现接口，使其支持以不同方式进行兼容分割

```go
// 声明类型 数据结构是string
type line string

// 对接口实现
func (l *line) Scan(state fmt.ScanState, verb rune) error {
	tok, err := state.Token(true, func(r rune) bool {
		return r != '-' // 以"-"分隔
	})
	if err != nil {
		return err
	}
	// 赋值操作
	*l = line(tok)
	return nil
}

var (
	s line
	n int
)
fmt.Sscanf("demo-125.txt", "%s-%d.txt", &s, &n) // 正常解析 s: demo
```

