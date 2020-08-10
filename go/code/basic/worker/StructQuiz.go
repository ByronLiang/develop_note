package worker

import (
    "fmt"
    "image/color"
    "log"
    "reflect"
    "test/tools"
    "unsafe"
)

type Dog struct {
    name string
    age int
    hobby string
    gender string
}

func (dog *Dog) Name() string {
    return dog.name
}

func (dog *Dog) SetName(name string) {
    dog.name = name
}

func ShowStruct()  {
    baby := Dog{
        name:  "tom",
        age:   0,
        hobby: "play",
    }
    //匿名结构体
    kk := struct {
        gender string
    }{
        gender: "male",
    }
    var owner *Dog
    //空结构体
    //owner = new(Dog)
    //owner = &Dog{"bob",21,"sleep"}
    //owner.name = "jk"
    owner = initData()
    owner.changeName("qq")
    //child := owner
    //child.name = "qq"
    fmt.Println(baby, kk, owner.name, unsafe.Sizeof(Dog{}))
}

//私有方法
func initData() *Dog {
    return &Dog{"lily",10,"eat","male"}
}

func (dog *Dog) changeName(name string)  {
    dog.name = name
    println(dog.name)
}

func MyStruct()  {
    none := struct {
        area string
    }{
        area: "abc",
    }
    cong := Dog{
        name:  "cong",
        age:   10,
        hobby: "play",
    }
    //同时拷贝内存地址与拷贝数据
    boy := &cong
    //新开辟内存地址拷贝数据
    //boy2 := cong
    boy.name = "kk"
    fmt.Println(cong.name, cong, &cong.name)
    fmt.Println(boy.name, boy, &boy.name)
    fmt.Println(none, &none.area)
}

func InitStruct()  {
    type MyArray [5]int
    type MyCap   []int
    var cc MyCap = make([]int, 5)
    type el interface {}
    var myEl el
    myEl = cc
    CheckType(myEl)
    switch myEl.(type) {
    case MyCap:
        fmt.Println("ss")
    }
    fmt.Println(cc)
}

func CheckType(el interface {})  {
    type MyCap   []int
    switch el.(type) {
        case MyCap:
            fmt.Println("ss")
    }
}

func ReflectTest()  {
    type test struct {
       Name     string      `json:"name" rules:"xx"`
       Age      int         `json:"age" rules:"ss"`
    }
    cc := test{Name:"jj", Age: 30}
    val := reflect.ValueOf(&cc).Elem()
    target := val.Field(0)
    fmt.Println("type of p:", target.Type())
    fmt.Println("is set ability of p:", target.CanSet())
    target.SetString("15")
    data, _ := tools.ToMap(&cc, "rules")
    fmt.Println(data)
}

type Point struct{ X, Y float64 }

type ColoredPoint struct {
    Point
    Color color.RGBA
}

type PColorPoint struct {
    *Point
    Color color.RGBA
}

func (p ColoredPoint) Distance(q *Point) float64 {
    return p.Point.Distance(*q)
}

func (point Point) Distance(q Point) float64 {
    return 2.11 + q.X
}

func MethodStruct()  {
    var c = ColoredPoint{
        Point: Point{1.2, 9.1},
        Color: color.RGBA{0, 0, 255, 255},
    }
    var d = PColorPoint{
        Point: &Point{1.1, 2.1},
        Color: color.RGBA{255, 0, 0, 255},
    }
    // 接收参数必须强一致性, 必须指向结构体的成员, 而不能是整个结构体
    log.Print(c.Distance(&c.Point))
    log.Print(d.Distance(*d.Point))

    // 以变量声明函数方法
    var mm func(point ColoredPoint, q *Point) float64
    mm = ColoredPoint.Distance
    kk := ColoredPoint.Distance
    // Distance实际上是指定了ColoredPoint对象为接收器的一个方法func (p ColoredPoint) Distance(q *Point)，
    // 但通过ColoredPoint.Distance得到的函数需要比实际的Distance方法多一个参数，
    // 即其需要用第一个额外参数指定接收器，后面排列Distance方法的参数。
    mm(c, &c.Point)
    kk(c, &c.Point)
}

