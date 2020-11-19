package singleflight

import (
    "sync"
)

type call struct {
    wg  sync.WaitGroup
    val interface{}
    err error
    match bool
}

type Group struct {
    mu sync.Mutex       // protects m
    m  map[string]*call // lazily initialized
    miss    map[string]struct{} // 存储空命中key
}

func (g *Group) Do(key string, fn func(key string) (interface{}, error, bool)) (interface{}, error) {
    g.mu.Lock()
    // 识别并动态初始化
    if g.m == nil {
        g.m = make(map[string]*call)
    }
    if g.miss == nil {
        g.miss = make(map[string]struct{})
    }
    if c, ok := g.m[key]; ok {
        g.mu.Unlock()
        c.wg.Wait()
        return c.val, c.err
    }
    c := new(call)
    c.wg.Add(1)
    g.m[key] = c
    g.mu.Unlock()

    if _, ok := g.miss[key]; ok {
        c.val, c.err = nil, nil
    } else {
        c.val, c.err, c.match = fn(key)
    }
    c.wg.Done()

    g.mu.Lock()
    if !c.match {
        g.miss[key] = struct{}{}
    }
    //delete(g.m, key)
    g.mu.Unlock()

    return c.val, c.err
}

func (g *Group) FindMissKeys() map[string]struct{} {
    return g.miss
}
