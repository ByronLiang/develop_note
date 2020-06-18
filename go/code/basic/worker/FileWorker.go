package worker

import (
    "encoding/json"
    "fmt"
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
