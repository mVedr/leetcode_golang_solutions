package main

import (
	"fmt"
	"sort"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type Node struct {
	ch  byte
	val int
}

func maxPointsInsideSquare(points [][]int, s string) int {
	arr := []Node{}
	for I, p := range points {
		arr = append(arr, Node{ch: s[I], val: max(abs(p[0]), abs(p[1]))})
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].val == arr[j].val {
			return arr[i].ch < arr[j].ch
		}
		return arr[i].val < arr[j].val
	})
	//fmt.Println(arr)
	ans := 0
	I := 0
	vis := make([]bool, 26)
	for I < len(arr) {
		J := I
		for J < len(arr) && arr[J].val == arr[I].val {
			if !vis[arr[J].ch-'a'] {
				vis[arr[J].ch-'a'] = true
				J++
			} else {
				return ans
			}
		}
		//fmt.Println("ans: ", ans)
		ans = J
		//	fmt.Println("ans: ", ans)
		I = J
	}
	return ans
}

func main() {
	fmt.Println(maxPointsInsideSquare([][]int{{2, 2}, {-1, -2}, {-4, 4}, {-3, 1}, {3, -3}}, "abdca"))
	fmt.Println(maxPointsInsideSquare([][]int{{1, 1}, {-2, -2}, {-2, 2}}, "aab"))
	fmt.Println(maxPointsInsideSquare([][]int{{1, 1}, {-1, -1}, {2, -2}}, "ccd"))
}
