package app

type Balance interface {
    DoBalance(resources []*Instance) (*Instance, error)
}
