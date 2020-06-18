package main

import (
    "fmt"
    "test/worker"
)

func main()  {
    items := map[string]string{
        "banana": "3",
        "apple": "2",
        "pear": "4",
        "lemon": "9",
        "pork": "20",
    }
    //line := worker.InitProductLine(items)
    line := worker.InitLine(items).(worker.ProductLine)
    //fmt.Println(line[0].Name)
    getItem(&line)
    fmt.Println("random get")
    buildCheck(&line, 2)
    fmt.Println("finished checked")
    worker.PrintAllItem(line)
}

func getItem(productLineCheckServiceImpl worker.ProductLineCheckService)  {
    item := productLineCheckServiceImpl.RandomGet().(*worker.Item)
    fmt.Println(item.Name, item.Num, item.IsChecked)
}

func buildCheck(productLineCheckServiceImpl worker.ProductLineCheckService, count int)  {
    productLineCheckServiceImpl.RandomCheck(count)
}
