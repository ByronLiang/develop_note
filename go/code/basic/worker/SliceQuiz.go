package worker

import "fmt"

func SliceTest() {
    //var blo []int

    var identifier = []int{1,3,2,0,0}
    l1 := len(identifier)
    aa := identifier[:len(identifier) - 2]
    identifier = identifier[:cap(identifier)]
    fmt.Println(l1)
    fmt.Println(aa)
    fmt.Println(identifier)
    fmt.Println(len(identifier))
}

func SliceQuiz()  {
    var no []int
    no = append(no, 1,2,3,4)
    blo := make([]int, 5)
    no1 := no[1:3]
    fmt.Println(no1, "", cap(no1))
    no1 = no1[0:3]
    fmt.Println(blo, cap(no1), no, no1)
}
