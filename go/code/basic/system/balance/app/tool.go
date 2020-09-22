package app

import (
    "fmt"
    "math/rand"
)

func InitBalanceData(len int, port int) (insets []*Instance) {
    for i := 0; i < len; i++ {
        host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
        one := NewInstance(host, port)
        insets = append(insets, one)
    }
    return
}

func InitBalanceResource(len int, port int) (resource Resources) {
    for i := 0; i < len; i++ {
        host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
        one := NewInstance(host, port)
        resource = append(resource, one)
    }
    return
}
