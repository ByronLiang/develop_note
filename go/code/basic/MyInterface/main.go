package main

import (
    "test/extra"
    "test/worker"
)

func main()  {
    //worker.InitAnimal()
    //worker.BasicInitInterfaceQuiz()
    //worker.BasicReflect()
    //worker.InterfaceCheck()
    testCheckInterface()

}

func testCheckInterface()  {
    //cat := new(extra.Cat)
    //cat.Name = "jimmy"
    //指针初始化
    dog := &extra.Dog{Name:"Gibber", Age:10}
    cat := &extra.Cat{Name:"jimmy", Age:22}
    worker.CheckAnimalInterface(cat)
    worker.CheckAnimalInterface(dog)
}
