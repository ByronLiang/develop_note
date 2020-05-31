package main

import "test/worker"

func main()  {
    //worker.ShowStruct()
    jk := worker.InitUser("jk", 20, "jk@gmail.com", "male")
    jk.ShowDetail()
    jk.ChangeName("super jk")
    jk.ShowDetail()
    //worker.MyStruct()
}
