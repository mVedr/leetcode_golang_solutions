package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func f(root *TreeNode, limit int, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		if sum+root.Val >= limit {
			return true
		}
		return false
	}
	left := f(root.Left, limit, sum+root.Val)
	right := f(root.Right, limit, sum+root.Val)

	if left == false {
		root.Left = nil
	}
	if right == false {
		root.Right = nil
	}
	if left || right {
		return true
	}
	return false
}

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	if root == nil {
		return nil
	}
	ans := f(root, limit, 0)
	if !ans {
		return nil
	}
	return root
}

func main() {

}
