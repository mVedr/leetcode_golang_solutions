package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		newR := &TreeNode{Val: val}
		newR.Left = root
		return newR
	}
	q := []*TreeNode{root}
	curr := 1
	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			p := q[0]
			q = q[1:]
			if curr == depth-1 {
				lc, rc := p.Left, p.Right
				p.Left = &TreeNode{Val: val}
				p.Right = &TreeNode{Val: val}
				p.Left.Left = lc
				p.Right.Right = rc
			} else {
				if p.Left != nil {
					q = append(q, p.Left)
				}
				if p.Right != nil {
					q = append(q, p.Right)
				}
			}
		}
		if curr == depth-1 {
			break
		}
		curr++
	}
	return root
}

func main() {

}
