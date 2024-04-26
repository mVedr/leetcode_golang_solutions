package main

import "fmt"

func canVisitAllRooms(g [][]int) bool {
	q := []int{}
	vis := make([]bool, len(g))
	q = append(q, 0)
	vis[0] = true
	for len(q) > 0 {
		n := len(q)
		for range n {
			tp := q[0]
			q = q[1:]
			for _, v := range g[tp] {
				if !vis[v] {
					vis[v] = true
					q = append(q, v)
				}
			}
		}
	}
	for _, val := range vis {
		if !val {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canVisitAllRooms([][]int{{1}, {2}, {3}, {}}))
	fmt.Println(canVisitAllRooms([][]int{{1, 3}, {3, 0, 1}, {2}, {0}}))
}
