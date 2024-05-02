package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func solve(root *TreeNode, fa *int) int {
	if root == nil {
		return 0
	}

	ans := 0
	sl, sr := solve(root.Left, fa), solve(root.Right, fa)
	(*fa) = max((*fa), sl)
	(*fa) = max((*fa), sr)
	(*fa) = max((*fa), root.Val+sl+sr)
	ans = max(ans, root.Val+sl)
	ans = max(ans, root.Val+sr)

	return max(ans, root.Val)
}

func maxPathSum(root *TreeNode) int {
	fa := math.MinInt32
	allNegative := true
	maxVal := math.MinInt32

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Val >= 0 {
			allNegative = false
		}
		maxVal = max(maxVal, node.Val)

		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)

	if allNegative {
		return maxVal
	}

	solve(root, &fa)
	return fa
}

func main() {

}
