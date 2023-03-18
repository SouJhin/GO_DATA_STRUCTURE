package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	arr := make([]int, 0)
	var preorder func(*TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		arr = append(arr, root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}
	preorder(root)
	return arr

}

func main() {
	t := TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 20,
			Right: &TreeNode{
				Val:   20,
				Left:  nil,
				Right: nil,
			},
			Left: nil,
		},
	}
	fmt.Printf("new(TreeNode{1,nil,nil}) =====> 🚀🚀🚀 %v\n", new(TreeNode))
	fmt.Printf("preorderTraversal() =====> 🚀🚀🚀 %v\n", preorderTraversal(&t))
}
