package main

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

func findBottomLeftValue(root *TreeNode) int {
	q := &Queue[*TreeNode]{}
	var ans int
	q.Push(root)
	for q.Len() > 0 {
		n := q.Len()
		fl := false
		for range n {
			tp := q.Pop()
			if !fl {
				fl = true
				ans = tp.Val
			}
			if tp.Left != nil {
				q.Push(tp.Left)
			}
			if tp.Right != nil {
				q.Push(tp.Right)
			}
		}
	}
	return ans
}

func main() {

}
