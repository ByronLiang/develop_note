package design

type ISingle interface {
    GetTitle() string
}

// 私有类型
type singleImpl struct {
    title    string
}

// 接口变量实例化
var SingleImpl ISingle

func (s singleImpl) GetTitle() string {
    return s.title
}

func NewSingleImpl(title string) ISingle {
    once.Do(func() {
        SingleImpl = &singleImpl{title: title}
    })
    if SingleImpl == nil {
        SingleImpl = &singleImpl{title: title}
    }
    return SingleImpl
}
