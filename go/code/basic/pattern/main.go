package main

import "test/pattern/aop"

func main()  {
    aopDev()
}

func aopDev()  {
    proxy := aop.NewUser("cc@gmail.com", "cc", "686123")
    proxy.Auth()
}
