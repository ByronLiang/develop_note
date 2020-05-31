package worker

import (
    "fmt"
    "unsafe"
)

type Dog struct {
    name string
    age int
    hobby string
    gender string
}

func ShowStruct()  {
    baby := Dog{
        name:  "tom",
        age:   0,
        hobby: "play",
    }
    //匿名结构体
    kk := struct {
        gender string
    }{
        gender: "male",
    }
    var owner *Dog
    //空结构体
    //owner = new(Dog)
    //owner = &Dog{"bob",21,"sleep"}
    //owner.name = "jk"
    owner = initData()
    owner.changeName("qq")
    //child := owner
    //child.name = "qq"
    fmt.Println(baby, kk, owner.name, unsafe.Sizeof(Dog{}))
}

//私有方法
func initData() *Dog {
    return &Dog{"lily",10,"eat","male"}
}

func (dog *Dog) changeName(name string)  {
    dog.name = name
    println(dog.name)
}

func MyStruct()  {
    none := struct {
        area string
    }{
        area: "abc",
    }
    cong := Dog{
        name:  "cong",
        age:   10,
        hobby: "play",
    }
    //同时拷贝内存地址与拷贝数据
    boy := &cong
    //新开辟内存地址拷贝数据
    //boy2 := cong
    boy.name = "kk"
    fmt.Println(cong.name, cong, &cong.name)
    fmt.Println(boy.name, boy, &boy.name)
    fmt.Println(none, &none.area)
}
