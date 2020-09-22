package app

import (
    "errors"
)

type RoundRobinBalance struct {
    curIndex int
}

func NewRoundRobinBalance(index int) *RoundRobinBalance {
    return &RoundRobinBalance{curIndex: index}
}

func (p *RoundRobinBalance) DoBalance(resources []*Instance) (resource *Instance, err error) {
    if len(resources) == 0 {
        resource = nil
        err = errors.New("No instance ")
        return
    }
    lens := len(resources)
    if p.curIndex >= lens {
        p.curIndex = 0
    }
    resource = resources[p.curIndex]
    p.curIndex = (p.curIndex + 1) % lens
    err = nil
    return
}
