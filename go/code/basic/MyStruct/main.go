package main

import (
    "test/extra"
    "test/worker"
)

func main()  {
    var person = extra.Person{}
    person.SetName("oo")
    name := person.Name()
    println(name)
    //worker.Basic()
    //worker.ShowStruct()
    jk := worker.InitUser("jk", 20, "jk@gmail.com", "male")
    jk.ShowDetail()
    jk.ChangeName("super jk")
    jk.ShowDetail()
    //worker.MyStruct()
}
