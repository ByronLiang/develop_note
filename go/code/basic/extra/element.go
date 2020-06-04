package extra

import "fmt"

type Num int

func (num *Num) PrettyNum() {
    fmt.Println(*num)
}
