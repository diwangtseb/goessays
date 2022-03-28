package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func toBinaryTree(items []interface{}) *TreeNode {
	n := len(items)
	if n == 0 {
		return nil
	}
	return inner(0, items)
}

func inner(index int, items []interface{}) *TreeNode {
	n := len(items)
	if n <= index || items[index] == nil {
		return nil
	}
	return &TreeNode{
		Val:   items[index].(int),
		Left:  inner(2*index+1, items),
		Right: inner(2*index+2, items),
	}
}

func main() {
	n := toBinaryTree([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9})
	println(n.Val)
}
