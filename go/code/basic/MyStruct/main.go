package main

import (
    "test/extra"
    "test/worker"
)

func main()  {
    //var person extra.Person
    var person = extra.Person{}
    person.SetName("oo")
    println(person.GetName())
    //worker.Basic()
    //worker.ShowStruct()
    jk := worker.NewUser("jk", 20, "jk@gmail.com", "male")
    jk.ShowDetail()
    jk.ChangeName("super jk")
    jk.ShowDetail()
    //worker.MyStruct()
}
