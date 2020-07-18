package main

import (
    "test/extra"
    "test/worker"
)

func main()  {
    //worker.Basic()
    //worker.ShowStruct()
    //worker.MyStruct()
    worker.ReflectTest()
}

func showUse()  {
    jk := worker.NewUser("jk", 20, "jk@gmail.com", "male")
    jk.ShowDetail()
    jk.ChangeName("super jk")
    jk.ShowDetail()
}

func initUser() {
    //var person extra.Person
    var person = extra.Person{}
    person.SetName("oo")
    println(person.GetName())
}
