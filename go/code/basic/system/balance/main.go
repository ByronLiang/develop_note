package main

import (
    "flag"
    "fmt"
    "test/system/balance/app"
)

var (
    factory *app.BalanceFactory
    port, len int
)

func init()  {
    factory = app.NewBalanceFactory()
    factory.AddContainer("robin", app.NewRoundRobinBalance(0))
    flag.IntVar(&port, "port", 8080, "set port")
    flag.IntVar(&len, "len", 4, "set init data length")
}

func main()  {
    flag.Parse()
    //hashResDev("/api/demo")
    commonBalDev()
}

func commonBalDev()  {
    data := app.InitBalanceData(len, port)
    for _, res := range data {
      fmt.Println("origin: ", res.String())
    }
    i := 3
    for i > 0 {
      ins, err := factory.DoBalance("robin", data)
      if err != nil {
          fmt.Println(err.Error())
      } else {
          fmt.Println(ins.String())
      }
      i--
    }
}

func hashResDev(key string)  {
    resources := app.InitBalanceResource(len, port)
    for _, resource := range resources {
        fmt.Println("resource: ", resource.String())
    }
    hash := app.NewHashBalance()
    resource, err := hash.DoBalance(resources, key)
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println(resource.String())
    }
}
