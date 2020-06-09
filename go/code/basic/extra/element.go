package extra

import "fmt"

type Num int

func (num *Num) PrettyNum() {
    fmt.Println(*num)
}

type Staff struct {
    manager
    Level   string
    Name    string
}

/**
不对外显示类 但类里的成员能对外访问
 */
type manager struct {
    // 成员能对外访问
    Name    string
    Email   string
}

type report interface {
    Detail() string
    Show()  string
}

/**
内层类实现方法
 */
func (staff manager) Detail() string {
    return "name is " + staff.Name
}

func (staff Staff) Show() string {
    return "My level is " + staff.Level + " my boss is " + staff.manager.Name +
        "my boss email is " + staff.Email
}

/**
Staff类同时实现接口方法
 */
func Tell(reportImpl report) string {
    return reportImpl.Detail() + " " + reportImpl.Show()
}
