package aop

import (
    "errors"
    "fmt"
)

type ValidateName struct {
}

func (ValidateName) Before(user *User) error {
    fmt.Println("ValidateName before")
    if user.Name == "admin" {
        return errors.New("admin can't be used")
    }

    return nil
}

func (ValidateName) After(user *User) {
    fmt.Println("ValidateName after")
    fmt.Printf("username:%s validate sucess\n", user.Name)
}

type ValidatePassword struct {
    MinLength   int
    MaxLength   int
}

// Before 前置校验
func (advice ValidatePassword) Before(user *User) error {
    fmt.Println("ValidatePassword before")
    if user.Pwd == "123456" {
        return errors.New("pass isn't strong")
    }

    if len(user.Pwd) > advice.MaxLength {
        return fmt.Errorf("len of pass must less than:%d", advice.MaxLength)
    }

    if len(user.Pwd) < advice.MinLength {
        return fmt.Errorf("len of pass must greater than:%d", advice.MinLength)
    }

    return nil
}

func (ValidatePassword) After(user *User) {
    fmt.Println("ValidatePassword after")
    fmt.Printf("password:%s validate sucess\n", user.Pwd)
}
