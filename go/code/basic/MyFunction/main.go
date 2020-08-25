package main

import (
    "test/worker"
)

var (
    lo1 = "kk"
    lo2 = false
)

func main()  {
    //worker.Quiz1()
    //lo1, lo2 = ss()
    //println(lo1, lo2)
    //worker.StringFunc()
    worker.MyMapQuiz()
    //worker.MapPoint()
    //worker.CoverString()
    worker.DiffEncodeByte()
}

func trans(double int, kk *int) {
    *kk = *kk * double
}

func ss() (kk string, k2 bool) {
    kk = "dd"
    k2 = true
    return
}
