package main

import (
	"fmt"
)

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

func parseBoolExpr(expression string) bool {
	st := Stack[rune]{}

	for _, ch := range expression {

		if ch == ')' {
			arr := []rune{}
			for st.Top() != '(' {
				arr = append(arr, st.Pop())
			}
			st.Pop()
			opt := st.Pop()
			if opt == '&' {
				ans := true
				for _, v := range arr {
					if v == 'f' {
						ans = false
						break
					}
				}
				if ans {
					st.Push('t')
				} else {
					st.Push('f')
				}
			} else if opt == '|' {
				ans := false
				for _, v := range arr {
					if v == 't' {
						ans = true
						break
					}
				}
				if ans {
					st.Push('t')
				} else {
					st.Push('f')
				}
			} else if opt == '!' {
				if arr[0] == 'f' {
					st.Push('t')
				} else {
					st.Push('f')
				}
			}
		} else {
			st.Push(ch)
		}
		//fmt.Printf("%c", st)
	}

	if st.Top() == 'f' {
		return false
	}
	return true
}

func main() {
	fmt.Println(parseBoolExpr("&(|(f))"))
	fmt.Println(parseBoolExpr("|(f,f,f,t)"))
	fmt.Println(parseBoolExpr("!(&(f,t))"))
}
