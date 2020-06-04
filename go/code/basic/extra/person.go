package extra

type Person struct {
    name string
    lastName  string
}

func (p *Person) Name() string {
    return p.name
}

func (p *Person) SetName(newName string) {
    (*p).name = newName
}
