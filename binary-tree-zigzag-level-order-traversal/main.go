package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := [][]int{}
	q := []*TreeNode{root}
	lvl := 0
	for len(q) > 0 {
		n := len(q)
		temp := []int{}
		for range n {
			tp := q[0]
			temp = append(temp, tp.Val)
			q = q[1:]
			if tp.Left != nil {
				q = append(q, tp.Left)
			}
			if tp.Right != nil {
				q = append(q, tp.Right)
			}
		}
		if lvl%2 == 1 {
			for I := range len(temp) / 2 {
				temp[I], temp[len(temp)-I-1] = temp[len(temp)-I-1], temp[I]
			}
		}
		lvl++
		ans = append(ans, temp)
	}
	return ans
}

func main() {

}
