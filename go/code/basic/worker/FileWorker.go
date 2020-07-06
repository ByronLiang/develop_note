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
