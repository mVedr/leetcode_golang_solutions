package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(1+maxDepth(root.Left), 1+maxDepth(root.Right))
}

func main() {

}
