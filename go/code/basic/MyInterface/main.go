package main

import (
    "fmt"
    "test/extra"
    "test/worker"
)

func main()  {
    //worker.InitAnimal()
    //worker.BasicInitInterfaceQuiz()
    //worker.BasicReflect()
    //worker.InterfaceCheck()
    testCheckInterface()
    //testStaff()
    //worker.TransJson()
}

func testCheckInterface()  {
    // 用于触发编译期的接口的合理性检查机制
    var _ extra.IAnimal = (*extra.Cat)(nil)
    //cat := new(extra.Cat)
    //cat.Name = "jimmy"
    //指针初始化
    dog := &extra.Dog{Name:"Gibber", Age:10}
    cat := &extra.Cat{Name:"jimmy", Age:22}
    worker.CheckAnimalInterface(cat)
    worker.CheckAnimalInterface(dog)
    worker.GetDogHabit(dog)
}

func testStaff()  {
    staff := &extra.Staff{Level:"basic"}
    staff.Name = "john"
    staff.Email = "john@gmail.com"
    // 不能透过非公开类对成员赋值
    //staff.manager.Name = "Ben"
    fmt.Println(extra.Tell(staff))
}
