package alog

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func InitListNode(init []int) *ListNode {
	head := &ListNode{}
	temp := head
	for _, val := range init {
		temp.Next = &ListNode{
			Val:  val,
			Next: nil,
		}
		temp = temp.Next
	}
	return head.Next
}

func ShowListNode(head *ListNode) {
	var data []int
	for head != nil {
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
				stack = stack[:(length - 1)]
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
		Val:  0,
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

/**
https://leetcode-cn.com/problems/remove-linked-list-elements/submissions/
*/
func RemoveElements(head *ListNode, target int) {
	for head != nil && head.Val == target {
		head = head.Next
	}
	current := head
	for current != nil && current.Next != nil {
		if current.Next.Val == target {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
}

func Partition(head *ListNode, x int) *ListNode {
	ltNode, geNode := &ListNode{
		Val:  0,
		Next: nil,
	}, &ListNode{
		Val:  0,
		Next: nil,
	}
	lt := ltNode
	ge := geNode
	for head != nil {
		if head.Val < x {
			lt.Next = head
			lt = lt.Next
		}
		if head.Val >= x {
			ge.Next = head
			ge = ge.Next
		}
		head = head.Next
	}
	ge.Next = nil
	lt.Next = geNode.Next
	return ltNode.Next
}

/**
https://leetcode-cn.com/problems/add-two-numbers/
*/
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	plus := 0
	newListNode := &ListNode{}
	current := newListNode
	for l1 != nil && l2 != nil {
		res := l1.Val + l2.Val + plus
		// 是否超出10 (是否需要进位)
		plus = res / 10
		// 取余数
		val := res % 10
		fmt.Printf("l1: %v; l2: %v; val: %v \n", l1.Val, l2.Val, val)
		current.Next = &ListNode{Val: val}
		current = current.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		res := l1.Val + plus
		// 是否超出10
		plus = res / 10
		// 取余数
		val := res % 10
		current.Next = &ListNode{Val: val}
		current = current.Next
		l1 = l1.Next
	}

	for l2 != nil {
		res := l2.Val + plus
		// 是否超出10
		plus = res / 10
		// 取余数
		val := res % 10
		current.Next = &ListNode{Val: val}
		current = current.Next
		l2 = l2.Next
	}

	// 检测是否仍有剩余位是否大于0, 大于0仍需进行进位处理
	if plus > 0 {
		current.Next = &ListNode{
			Val: plus,
		}
		current = current.Next
	}

	return newListNode.Next
}

/**
递归合并两序列
https://leetcode-cn.com/problems/merge-two-sorted-lists/
*/
func MergeTwoListsWithLoop(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = MergeTwoListsWithLoop(l1.Next, l2)
		return l1
	} else {
		l2.Next = MergeTwoListsWithLoop(l1, l2.Next)
		return l2
	}
}

/**
非递归合并链表
*/
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newListNode := new(ListNode)
	current := newListNode
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}
	if l1 == nil {
		current.Next = l2
	}
	if l2 == nil {
		current.Next = l1
	}
	return newListNode.Next
}
