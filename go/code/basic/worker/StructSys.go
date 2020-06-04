package worker

import (
    "fmt"
    "time"
)

/**
user 继承部分person 结构体
 */
type user struct {
    name    string
    age     int
    email   string
    person
}

type person struct {
    gender  string
    hobby   string
}

type age int

type myTime struct {
    time.Time //anonymous field
}

func InitUser(name string, age int, email string, gender string) user {
    return user{name, age, email,person{gender,""}}
}

/**
Time 拓展 package:time 里没有的方法
 */
func (t myTime) first3Chars() string {
    return t.Time.String()[0:3]
}

func Basic()  {
    /**
    类型和作用在它上面定义的方法必须在同一个包里定义
     */
    //ss := new(time.Time)
    //ss.first3Chars()
    m := myTime{time.Now()}
    // 调用匿名Time上的String方法
    fmt.Println("Full time now:", m.String())
    // 调用myTime.first3Chars
    fmt.Println("First 3 chars:", m.first3Chars())
    fmt.Println(age(20))
}

func (user user) ShowDetail()  {
    fmt.Printf(
        "name: %s age: %d gender: %s hobby: %s \n",
        user.name, user.age, user.gender, user.hobby)
}

func (user *user) ChangeName(name string) {
    user.name = name
}

func (user *user) ChangeEmail(email string) {
    user.email = email
}
func (user *user) SetHobby(hobby string) {
    user.person.hobby = hobby
}
