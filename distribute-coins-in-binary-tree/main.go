package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	ans = 0
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func dfs(root *TreeNode) int {
	if root.Left != nil {
		root.Val += dfs(root.Left)
	}
	if root.Right != nil {
		root.Val += dfs(root.Right)
	}
	tmp := root.Val - 1
	ans += abs(tmp)
	return tmp
}

func distributeCoins(root *TreeNode) int {
	ans = 0
	dfs(root)
	return ans
}

func main() {

}
