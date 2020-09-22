package app

import (
    "fmt"
    "sync"
)

type BalanceFactory struct {
    sync.RWMutex
    containers  map[string]Balance
}

func NewBalanceFactory() *BalanceFactory {
    return &BalanceFactory{containers: make(map[string]Balance)}
}

func (b BalanceFactory) AddContainer(key string, balance Balance)  {
    b.Lock()
    defer b.Unlock()
    b.containers[key] = balance
}

func (b BalanceFactory) DelContainer(key string)  {
    b.Lock()
    defer b.Unlock()
    delete(b.containers, key)
}

func (b BalanceFactory) GetContainer(key string) (balance Balance, err error) {
    b.RLock()
    defer b.RUnlock()
    balance, ok := b.containers[key]
    if !ok {
        err = fmt.Errorf(" None %s balancer ", key)
    }
    return
}

func (b BalanceFactory) DoBalance(key string, ins []*Instance) (*Instance, error) {
    c, err := b.GetContainer(key)
    if err != nil {
        return nil, err
    }
    return c.DoBalance(ins)
}
