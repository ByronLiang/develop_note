package alog

import (
    "fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func InitListNode(init []int) *ListNode {
    head := &ListNode{
        Val: 1,
        Next: nil,
    }
    temp := head
    for _, val := range init {
        temp.Next = &ListNode{
            Val: val,
            Next: nil,
        }
        temp = temp.Next
    }
    return head
}

func ShowListNode(head *ListNode)  {
    var data []int
    for head != nil  {
        data = append(data, head.Val)
        head = head.Next
    }
    fmt.Println(data)
}

/**
回文链表
https://leetcode-cn.com/problems/palindrome-linked-list/
 */
func IsPalindrome(head *ListNode) bool {
    var stack []int
    fast := head
    slow := head
    for fast != nil && fast.Next != nil {
        stack = append(stack, slow.Val)
        fast = fast.Next.Next
        slow = slow.Next
    }
    fmt.Println(slow.Val, stack)
    // 无法通过stack的length 判断链表长度是奇数还是偶数
    // 判断原链表长度是奇数 需要对中间点进行进一，再进行回文判断
    // 偶数长度链表 直接进行回文判断
    if fast != nil {
        fmt.Println("skip the middle point")
        slow = slow.Next
    }
    for slow != nil {
        length := len(stack)
        if length > 0 && stack != nil {
            i := stack[length-1]
            if i != slow.Val {
                fmt.Println("none match")
                return false
            } else {
                // 出栈
                stack = stack[:(length-1)]
            }
        } else {
            fmt.Println("none match")
            return false
        }
        slow = slow.Next
    }
    return true
}

/**
https://leetcode-cn.com/problems/swap-nodes-in-pairs/submissions/
 */
func SwapPairs(head *ListNode) *ListNode {
    NewNode := &ListNode{
        Val: 0,
        Next: nil,
    }
    NewNode.Next = head
    current := NewNode
    //if current.Next == nil {
    //    return nil
    //}
    //if current.Next.Next == nil {
    //    return current.Next
    //}
    for current.Next != nil && current.Next.Next != nil {
        start := current.Next
        end := current.Next.Next
        // 将NewNode节点拼接起来
        current.Next = end
        start.Next = end.Next
        end.Next = start
        current = start
    }
    return NewNode.Next
}
