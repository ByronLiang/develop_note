package worker

import (
    "fmt"
    "reflect"
    "test/extra"
)

func InitAnimal()  {
    var animalTarget extra.IAnimal
    // 以Cat类型创建一个值
    //var tom extra.Cat
    //tom := extra.Cat{}
    // 以指针创建对象
    tom := new(extra.Cat)
    tom.Name = "kk"
    tom.AddHobby("play game")
    extra.ChangHobby(tom, "play ball")
    fmt.Println(tom.Habit(), tom.Attack(), tom.Hobby)

    //IAnimal接口是基于Cat结构体的接口实现
    animalTarget = tom
    /**
      switch-type 对接口类型用法
    */
    switch animalTarget.(type) {
        case *extra.Cat:
            fmt.Println("dd")
        default:
            fmt.Println("qq")
    }

    //检验接口实现与具体结构体的识别
    res, ok := animalTarget.(*extra.Cat)
    fmt.Println(res, ok)
    var animal = extra.IAnimal(tom)
    //var animal extra.IAnimal = extra.Cat{"Tom",23}
    fmt.Println(animal.Habit(), tom.Attack())
}

func BasicInitInterfaceQuiz()  {
    //num := extra.Num(32)
    //指定对某类型进行初始化
    var num extra.Num = 32
    (&num).PrettyNum()

}

/**
全局声明
 */
type (
    MyArray [5]int
    MyCap   []int
)
func InterfaceCheck()  {
    kk := MyCap(make([]int, 5))
    var cc MyCap = make([]int, 5)
    //type el interface {}
    var myEl interface {}
    myEl = cc
    checkType(kk)
    checkType(myEl)
    fmt.Println(cc)
}

func checkType(el interface {})  {
    switch el.(type) {
    case MyCap:
        fmt.Println("ss")
    }
}

func BasicReflect()  {
    type T struct {
        A int
        B string
    }
    type kk int
    //var t = T{12, "tom"}
    //初始化接口类型的值
    var ll interface{} = 32
    mm := kk(32)
    fmt.Println(reflect.TypeOf(ll), reflect.ValueOf(ll), reflect.ValueOf(mm))
}

func CheckAnimalInterface(animal interface{})  {
    if pet,ok := animal.(extra.IAnimal); ok {
        extra.RecommendSelf(pet)
    } else {
        fmt.Println("No implementation IAnimal interface")
    }
}

// 即便Dog结构体未完全实现animal 接口 animal作为空接口 接收对象
// animal.(*extra.Dog) 换取回Dog 结构体对象
func GetDogHabit(animal interface{})  {
    if pet, ok := animal.(*extra.Dog); ok {
        fmt.Println(pet.Habit())
    } else {
        fmt.Println("No implementation IAnimal interface")
    }
}
