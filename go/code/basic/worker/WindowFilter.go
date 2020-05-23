package worker

import "fmt"

func WindowMax(target []int, size int) int {
    var (
        window []int
        //k = 0
        max = 0
    )
    for num, _ := range target {
        println(num, target[num])
    }
    
    for i := 0; i < len(target); i ++ {
        if i < size - 1 {
            window = append(window, target[i])
        } else {
            window = append(window, target[i])
            fmt.Print("start")
            fmt.Println(window)
            res :=MultiAdd(window)
            if res > max {
                max = res
            }
            window = append(window[:0], window[1:]...)
            fmt.Print("end")
            fmt.Println(window)
        }
    }
    return max
}

func MultiAdd(target []int) int {
    var res = 0
    for _,value := range target {
        res += value
    }
    return res
}
