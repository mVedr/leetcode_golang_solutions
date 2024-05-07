package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue[T any] struct {
	arr []T
}

func (q *Queue[T]) Push(val T) {
	q.arr = append(q.arr, val)
}

func (q *Queue[T]) Len() int {
	return len(q.arr)
}

func (q *Queue[T]) Pop() T {
	tp := q.Top()
	q.arr = q.arr[1:len(q.arr)]
	return tp
}

func (q *Queue[T]) Top() T {
	return q.arr[0]
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	if k == 0 {
		return []int{target.Val}
	}
	g := make(map[int][]int)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root.Left != nil {
			g[root.Val] = append(g[root.Val], root.Left.Val)
			g[root.Left.Val] = append(g[root.Left.Val], root.Val)
			dfs(root.Left)
		}
		if root.Right != nil {
			g[root.Val] = append(g[root.Val], root.Right.Val)
			g[root.Right.Val] = append(g[root.Right.Val], root.Val)
			dfs(root.Right)
		}
	}
	dfs(root)
	fmt.Println(g)
	q := &Queue[int]{}
	vis := make(map[int]bool)
	vis[target.Val] = true
	q.Push(target.Val)
	cnt := 0
	for q.Len() > 0 {
		n := q.Len()
		ans := []int{}
		for i := 0; i < n; i++ {
			tp := q.Pop()

			for _, neigh := range g[tp] {
				if !vis[neigh] {
					vis[neigh] = true
					if cnt+1 == k {
						ans = append(ans, neigh)
					}
					q.Push(neigh)

				}
			}
		}
		if len(ans) > 0 {
			return ans
		}
		cnt++
	}
	return []int{}
}

func main() {

}
