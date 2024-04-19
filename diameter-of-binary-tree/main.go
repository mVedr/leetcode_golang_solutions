package main

var (
	ans = 0
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func solve(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	lh := 1 + solve(root.Left, ans)
	rh := 1 + solve(root.Right, ans)
	*ans = max(*ans, lh+rh-2)
	return max(lh, rh)
}

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	solve(root, &ans)
	return ans
}

func main() {

}
