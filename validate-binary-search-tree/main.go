package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBSTCheck(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}
	ans := true
	if root.Left != nil {
		ans = ans && (root.Left.Val < root.Val)
	}
	if root.Right != nil {
		ans = ans && (root.Right.Val > root.Val)
	}
	return ans && isValidBSTCheck(root.Left, min, root.Val) && isValidBSTCheck(root.Right, root.Val, max)
}

func isValidBST(root *TreeNode) bool {
	return isValidBSTCheck(root, math.MinInt, math.MaxInt)
}

func main() {

}
