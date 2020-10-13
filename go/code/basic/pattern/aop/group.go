package aop

// UserAopGroup 通知管理组
type UserAopGroup struct {
    items []UserAop
}

// Add 注入可选通知
func (g *UserAopGroup) Add(aop UserAop) {
    g.items = append(g.items, aop)
}

func (g *UserAopGroup) Before(user *User) error {
    for _, item := range g.items {
        if err := item.Before(user); err != nil {
            return err
        }
    }

    return nil
}

// After
func (g *UserAopGroup) After(user *User) {
    for _, item := range g.items {
        item.After(user)
    }
}
