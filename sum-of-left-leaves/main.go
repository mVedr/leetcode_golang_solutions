package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*TreeNode{root}
	ans := 0
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node.Left != nil {
			if node.Left.Left == nil && node.Left.Right == nil {
				ans += node.Left.Val
			} else {
				q = append(q, node.Left)
			}
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}
	return ans
}

func main() {

}
