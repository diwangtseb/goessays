package ds

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (tn *TreeNode) Bfs() {
	if tn == nil {
		return
	}
	queue := []*TreeNode{tn}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Println(node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

}

func (tn *TreeNode) Dfs() {
	if tn == nil {
		return
	}
	fmt.Println(tn.Val)
	tn.Left.Dfs()
	tn.Right.Dfs()
}

func (tn *TreeNode) Height() int {
	if tn == nil {
		return 0
	}
	leftH := tn.Left.Height()
	rightH := tn.Right.Height()
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	return max(leftH, rightH) + 1
}

func recursionFunc(n *TreeNode, val int) *TreeNode {
	if n == nil {
		return nil
	}
	if n.Val > val {
		if n.Left == nil {
			n.Left = &TreeNode{Val: val}
		} else {
			n.Left = recursionFunc(n.Left, val)
		}
	} else {
		if n.Right == nil {
			n.Right = &TreeNode{Val: val}
		} else {
			n.Right = recursionFunc(n.Right, val)
		}
	}
	return n
}

func InsertBstByRecursionArray(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: arr[0],
	}
	for i := 1; i < len(arr); i++ {
		recursionFunc(root, arr[i])
	}
	return root
}

func InsertBbtByRecursionArray(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	mid := (len(arr) - 1) / 2
	return &TreeNode{
		Val:   arr[mid],
		Left:  InsertBbtByRecursionArray(arr[:mid]),
		Right: InsertBbtByRecursionArray(arr[mid+1:]),
	}
}

func StackTraverse(root *TreeNode) []int {
	rsp := []int{}
	if root == nil {
		return rsp
	}
	stack := []*TreeNode{}
	current := root
	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		pos := len(stack) - 1
		current = stack[pos]
		stack = stack[:pos]

		rsp = append(rsp, current.Val)

		current = current.Right
	}
	return rsp
}
