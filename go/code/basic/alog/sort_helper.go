package alog

import "fmt"

type SortHelper struct {

}

func (sh SortHelper) MergeVSort(data []int) []int {
    if len(data) < 2 {
        return data
    }
    mid := len(data) >> 1
    left := sh.MergeVSort(data[0:mid])
    fmt.Println("mid", mid, left, data)
    right := sh.MergeVSort(data[mid:])
    fmt.Println(left, right)
    return bing(left, right)
}

func bing(left, right []int) []int {
    var res []int
    j := 0
    k := 0
    for j < len(left) && k < len(right) {
        if left[j] > right[k] {
            res = append(res, right[k])
            k++
        } else {
            res = append(res, left[j])
            j++
        }
    }
    res = append(res, left[j:]...)
    res = append(res, right[k:]...)
    return res
}

func (sh SortHelper) MergeSort(data []int, start, end int)  {
    if end - start > 1 {
        mid := start + (end - start + 1) >> 1
        sh.MergeSort(data, start, mid)
        sh.MergeSort(data, mid, end)
        fmt.Println(start, mid, end)
        merge(data, start, mid, end)
    }
}

func (sh SortHelper) QuickSort(data []int)  {
    quickSort(data, 0, len(data) - 1)
}

func quickSort(data []int, l, r int)  {
    if l < r {
        mid := quickSortPartition(data, l, r)
        fmt.Println("mid: ", mid, l, r)
        quickSort(data, l, mid - 1)
        quickSort(data, mid + 1, r)
    }
}

func (sh SortHelper) TopQuickSort(data []int, top int, isDesc bool) []int {
    if top > len(data) {
        return nil
    }
    // 最大Top k
    if isDesc {
        topQuickSort(data, 0, len(data) - 1, len(data) - top)
        return data[(len(data) - top):]
    }
    // 最小Top k
    topQuickSort(data, 0, len(data) - 1, top - 1)
    return data[:top]
}

func topQuickSort(data []int, l, r, top int)  {
    if l < r {
        mid := quickSortPartition(data, l, r)
        fmt.Println("mid: ", mid, l, r, data)
        if top > mid {
            topQuickSort(data, mid + 1, r, top)
        } else if top < mid {
            topQuickSort(data, l, mid - 1, top)
        } else if top == mid {
            return
        }
    }
}

func quickSortPartition(data []int, l, r int) int {
    target := data[l]
    for l < r {
        for l < r && data[r] >= target {
            r--
        }
        data[l], data[r] = data[r], data[l]
        for l < r && data[l] <= target {
            l++
        }
        data[l], data[r] = data[r], data[l]
    }
    return l
}

func merge(data []int, start, mid, end int)  {
    lSize := mid - start
    rSize := end - mid
    s := end - start
    temp := make([]int, 0, s)
    l, r := 0, 0
    for l < lSize && r < rSize {
        lValue := data[start+l]
        rValue := data[mid+r]
        if lValue < rValue {
            temp = append(temp, lValue)
            l++
        } else {
            temp = append(temp, rValue)
            r++
        }
    }
    if l < lSize {
        temp = append(temp, data[start+l:mid]...)
    }
    if r < rSize {
        temp = append(temp, data[mid+r:end]...)
    }

    for i:=0; i < s; i++ {
        data[start+i] = temp[i]
    }
    fmt.Println("temp: ", temp, data)
}
