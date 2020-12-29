package alog

import (
    "testing"
)

func TestSortHelper_MergeSort(t *testing.T) {
    sh := SortHelper{}
    data := []int{8, 3, 1, 9, 2}
    sh.MergeSort(data, 0, len(data))
    t.Log(data)
}

func TestSortHelper_MergeVSort(t *testing.T) {
    sh := SortHelper{}
    data := []int{8, 3, 1, 9, 2}
    res := sh.MergeVSort(data)
    t.Log(res)
}

func TestSortHelper_QuickSort(t *testing.T) {
    sh := SortHelper{}
    data := []int{8, 3, 1, 9, 2, 5}
    sh.QuickSort(data)
    t.Log(data)
}

func TestSortHelper_TopQuickSort(t *testing.T) {
    sh := SortHelper{}
    data := []int{8, 3, 1, 9, 2, 5}
    desc := sh.TopQuickSort(data, 3, true)
    asc := sh.TopQuickSort(data, 2, false)
    t.Log(data)
    t.Log(desc)
    t.Log(asc)
}
