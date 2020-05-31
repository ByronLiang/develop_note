package worker

import "fmt"

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

func InitUser(name string, age int, email string, gender string) user {
    return user{name, age, email,person{gender,""}}
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
