package alog

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func newNode(val int) *TreeNode {
    return &TreeNode{Val: val}
}

func InitTreeData() (root *TreeNode) {
    root = &TreeNode{
        Val: 4,
    }
    root.Left = newNode(9)
    root.Right = newNode(0)
    root.Left.Left = newNode(5)
    root.Left.Right = newNode(1)
    return
}

func SumNumbers(root *TreeNode) int {
    return dfs(root, 0)
}

func dfs(root *TreeNode, sum int) int {
    if root == nil {
        return 0
    }
    sum = sum*10 + root.Val
    if root.Left == nil && root.Right == nil {
        return sum
    }
    return dfs(root.Left, sum) + dfs(root.Right, sum)
}
