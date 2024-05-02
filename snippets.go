package main

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
