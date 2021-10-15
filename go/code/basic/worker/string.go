package worker

import (
	"fmt"
	"sort"
	"strings"
)

func StringFunc() {
	var (
		index int
		word  string
	)
	paper := []string{"today", "morning"}
	//查询字符[字节查询]
	index, word = GetByte(paper[0], 166)
	//拼接字符串
	data := strings.Join(paper, ",")
	//查询后缀
	res := strings.HasSuffix("2010-11-10.jpg", ".jpg")
	fmt.Println(data, res, paper[0][0], index, word)
}

/**
函数参数
数据结构需要与形参一致
*/
func GetByteIndex(target rune, ss int) bool {
	return target == 116 || target == 161
}

func GetByte(words string, findByte rune) (index int, word string) {
	index = strings.IndexFunc(words, func(target rune) bool {
		return target == findByte
	})
	if index >= 0 {
		word = string(words[index])
	} else {
		word = "unfold"
	}
	return
}

func CoverString() {
	list := []string{"banner", "apple", "country"}
	sort.Strings(list)
	fmt.Println(list)
}

func DiffEncodeByte() {
	var (
		// 8位 满足一个英文字符
		en = []byte("good-day")
		// byte 接收中文字符 会出现乱码
		// rune 32位 四字符 能完整接收中/英文字符(3位字符)
		cn = []rune("你好-明天, What I Can Do")
	)
	enTxt := en[0:4]
	cnTxt := cn[0:11]
	fmt.Println(string(enTxt), string(cnTxt))
}
