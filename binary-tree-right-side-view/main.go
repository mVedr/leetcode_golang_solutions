package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	q := []*TreeNode{root}
	ans := []int{}
	for len(q) > 0 {
		n := len(q)
		tmp := -1
		for _ = range n {
			tp := q[0]
			q = q[1:]
			tmp = tp.Val
			if tp.Left != nil {
				q = append(q, tp.Left)
			}
			if tp.Right != nil {
				q = append(q, tp.Right)
			}

		}
		ans = append(ans, tmp)
	}
	return ans
}

func main() {

}
