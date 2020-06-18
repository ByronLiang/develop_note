package worker

import (
    "fmt"
    "math/rand"
)

type Item struct {
    Name        string
    Note        string
    Num         int
    IsChecked   bool
}

type ProductLine []*Item

type ProductLineCheckService interface {

    RandomGet() interface{}

    RandomCheck(count int)
}

func InitProductLine(products map[string]string) ProductLine {
    line := make(ProductLine, len(products))
    i := 0
    for name, note := range products {
        line[i] = &Item{
            Name:      name,
            Note:      note,
            Num:       i,
            IsChecked: false,
        }
        i ++
    }
    return line
}

func InitLine(products map[string]string) interface{} {
    line := make(ProductLine, len(products))
    i := 0
    for name, note := range products {
        line[i] = &Item{
            Name:      name,
            Note:      note,
            Num:       i,
            IsChecked: false,
        }
        i ++
    }
    return line
}

func (p *ProductLine) RandomGet() interface{} {
    product := *p
    num := rand.Intn(len(product))
    return product[num]
}

func (p *ProductLine) RandomCheck (count int)  {
    product := *p
    for i := 0; i < count; i ++ {
        num := rand.Intn(len(product))
        item := (*p)[num]
        //if item.Num >= i {
        //    item.IsChecked = true
        //}
        item.IsChecked = true
    }
}

func PrintAllItem(p ProductLine)  {
    for _, item := range p {
        fmt.Println(item.Name, item.Num, item.IsChecked)
    }
}


