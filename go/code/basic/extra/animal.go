package extra

import "fmt"

type IAnimal interface {
    Habit() string
    Sound() string
    AddHobby(hobby string)
}

type Cat struct {
    Name    string
    Age     int
    Hobby   []string
}

type Dog struct {
    Name    string
    Age     int
}

func (dog Dog) Habit() string {
    return dog.Name + " like to eat bone"
}

func (cat Cat) Habit() string {
    return cat.Name + " like to eat fish"
}

func (cat Cat) Sound() string {
    return "miaomiao"
}

func (cat *Cat) AddHobby(hobby string)  {
    (*cat).Hobby = append((*cat).Hobby, hobby)
}

func (cat Cat) Attack () string {
    return "attack by hand"
}

func ChangHobby(obj IAnimal, hobby string)  {
    /**
      需要指针传递
    */
    obj.AddHobby(hobby)
}

func RecommendSelf(obj IAnimal) {
    fmt.Println("my hobby is "+ obj.Habit() + " my sound is " + obj.Sound())
}



