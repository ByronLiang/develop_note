package alog

import "testing"

func TestAddTwoNumbers(t *testing.T) {
    l1 := InitListNode([]int{9,9,9,9,9,9,9})
    l2 := InitListNode([]int{9,9,9,9})
    res := AddTwoNumbers(l1, l2)
    ShowListNode(res)
}

func TestSwapPairs(t *testing.T) {
    node := InitListNode([]int{2,3,4})
    ShowListNode(node)
    newNode := SwapPairs(node)
    ShowListNode(newNode)
}

func TestIsPalindrome(t *testing.T) {
    node := InitListNode([]int{1,2,2,1})
    IsPalindrome(node)
}
