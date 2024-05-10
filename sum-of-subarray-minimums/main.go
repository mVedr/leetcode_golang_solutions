package main

import "fmt"

const MOD = 1000000007

type Node struct {
	val int
	ind int
}

type Stack[T any] struct {
	arr []T
}

func (s *Stack[T]) Push(val T) {
	s.arr = append(s.arr, val)
}

func (s *Stack[T]) Len() int {
	return len(s.arr)
}

func (s *Stack[T]) Pop() T {
	tp := s.Top()
	s.arr = s.arr[:len(s.arr)-1]
	return tp
}

func (s *Stack[T]) Top() T {
	return s.arr[len(s.arr)-1]
}

func sumSubarrayMins(arr []int) int {
	n := len(arr)
	ls, rs := make([]int, n), make([]int, n)
	stk := &Stack[int]{}

	for i := 0; i < n; i++ {
		for stk.Len() > 0 && arr[stk.Top()] >= arr[i] {
			stk.Pop()
		}
		if stk.Len() == 0 {
			ls[i] = -1
		} else {
			ls[i] = stk.Top()
		}
		stk.Push(i)
	}

	stk = &Stack[int]{}

	for i := n - 1; i >= 0; i-- {
		for stk.Len() > 0 && arr[stk.Top()] > arr[i] {
			stk.Pop()
		}
		if stk.Len() == 0 {
			rs[i] = n
		} else {
			rs[i] = stk.Top()
		}
		stk.Push(i)
	}

	ans := 0
	for i := 0; i < n; i++ {
		l, r := i-ls[i], rs[i]-i
		ans = (ans + (arr[i]*l*r)%MOD) % MOD
	}

	return ans
}

func main() {
	fmt.Println(sumSubarrayMins([]int{3, 1, 2, 4}))
	fmt.Println(sumSubarrayMins([]int{11, 81, 94, 43, 3}))
}
