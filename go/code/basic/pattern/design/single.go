package design

import (
    "sync"
)

type Single struct {
    Name    string
}

var (
    once sync.Once
    single *Single
)

/**
单例模式
 */
func Init(name string) *Single {
    once.Do(func() {
        single = &Single{Name: name}
    })
    if single == nil {
        single = &Single{Name: name}
    }
    return single
}

func (sin *Single) GetName() string {
    return sin.Name
}
