package main

type Stack struct {
	arr []any
}

func (s *Stack) Push(val any) {
	s.arr = append(s.arr, val)
}

func (s *Stack) Len() int {
	return len(s.arr)
}

func (s *Stack) Pop() any {
	tp := s.Top()
	s.arr = s.arr[:len(s.arr)-1]
	return tp
}

func (s *Stack) Top() any {
	return s.arr[len(s.arr)-1]
}

type Queue struct {
	arr []any
}

func (q *Queue) Push(val any) {
	q.arr = append(q.arr, val)
}

func (q *Queue) Len() int {
	return len(q.arr)
}

func (q *Queue) Pop() any {
	tp := q.Top()
	q.arr = q.arr[1:len(q.arr)]
	return tp
}

func (q *Queue) Top() any {
	return q.arr[0]
}
