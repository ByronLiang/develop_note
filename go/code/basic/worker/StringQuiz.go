package worker

import (
    "fmt"
    "strings"
)

var (
    index int
    word string
)

func StringFunc()  {
    paper := []string{"today","morning"}
    //查询字符
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
