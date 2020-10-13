package aop

import "fmt"

type User struct {
    Account     string
    Pwd         string
    Name        string
}

type UserAop interface {

    // Before 前置通知
    Before(user *User) error

    // After 后置通知
    After(user *User)
}

// Auth 验证
func (u *User) Auth() {
    // 实际业务逻辑
    fmt.Printf("register account: %s, name: %s\n", u.Account, u.Name)
}
