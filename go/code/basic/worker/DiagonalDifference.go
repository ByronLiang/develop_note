package worker

import (
    "fmt"
    "math"
)

/**
计算对角差
 */
func CalDiagonalDifference()  {
    var (
        test [][]int
        line int
        opt int
    )
    test = [][]int{{3}, {11, 2, 4}, {4, 5, 6}, {10, 8, -12}}
    N := test[0][0]
    for i, j := 1, 0 ; i <= N && j < N; i, j = i+1, j+1 {
       line += test[i][j]
       opt += test[i][(N-i)]
    }

    fmt.Println(line, opt, int32(math.Abs(float64(line - opt))))
}
