package worker

import (
    "bufio"
    "bytes"
    "encoding/json"
    "fmt"
    "os"
    "strconv"
)

type Person struct {
    Name    string
    Age     int
    Parents []string
    Hobby   []string
}

func TransJson()  {
    var data []Person
    text := `[{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}]`
    target := []byte(text)
    err := json.Unmarshal(target, &data)
    if err == nil {
        for _, item := range data {
            word := "name is " + item.Name + " age is " + strconv.Itoa(item.Age)
            if item.Hobby == nil {
                word += " No hobby"
            }
            word += " parents is "
            for _, parent := range item.Parents {
                word += parent + " "
            }
            fmt.Println(word)
        }
        //fmt.Println(data)
    }
    //fmt.Println(error)
}

func EncodeJsonSample()  {
    params := Person{
        "john",
        20,
        []string{"kelly", "Tom"},
        nil}
    byteBuf := bytes.NewBuffer([]byte{})
    encoder := json.NewEncoder(byteBuf)
    // 特殊字符不转义
    encoder.SetEscapeHTML(false)
    err := encoder.Encode(params)
    // 流程相同
    //res, _ := json.Marshal(params)
    //fmt.Println(string(res))

    if err != nil {
        panic(err)
    }
    data := byteBuf.String()
    fmt.Println(data)
}

func TestInput() {
    fmt.Println("start")
    counts := make(map[string]int)
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        counts[input.Text()]++
        if input.Text() == "end" {
            break
        }
    }
    // NOTE: ignoring potential errors from input.Err()
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}

func ByteBuff() {
    var readInput *bufio.Reader
    readInput = bufio.NewReader(os.Stdin)
    fmt.Println("Input your name ?")
    name, _ := readInput.ReadString('\n')
    fmt.Printf("your name is %s", name)
}

type ArticlePoint struct {
    Id      string          `json:"id"`
    Name    *string         `json:"name,omitempty"`
    Desc    *string         `json:"desc,omitempty"`
    View    int             `json:"view,omitempty"`
}

type Article struct {
    Id      string          `json:"id"`
    Name    string          `json:"name,omitempty"`
    Desc    string          `json:"desc,omitempty"`
    View    int             `json:"view,omitempty"`
}

func TransObj()  {
    // 非指针数据类型具备固有的空安全性; string/int 默认为空字符串/0显示
    jsonData := `{"id":"1234","name":"xyz"}`
    //jsonData := `{"id":"1234","name":"xyz","desc":""}`
    req := Article{}
    _ = json.Unmarshal([]byte(jsonData), &req)
    fmt.Println(req.Name, req.View, req.Desc)
}

func TransObjPoint()  {
    demo := `{"id":"1234","name":"xyz"}`
    //jsonData := `{"id":"1234","name":"xyz","desc":""}`
    req := ArticlePoint{}
    _ = json.Unmarshal([]byte(demo), &req)
    fmt.Println(*req.Name, req.View)
    // 指针数据类型, 需要进行空安全性检测: 判断是否空指针
    // 定义了指针，则这些数据类型在未手动设置的情况下默认为空
    // 不验证可空性的情况下访问那些指针的数据可能会导致应用程序崩溃
    if req.Desc != nil {
        fmt.Println("desc: ", *req.Desc)
    }
}

