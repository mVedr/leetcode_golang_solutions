package main

import (
	"fmt"
)

func dist(src, n int, g map[int][]int) []int {
	q := []int{}
	vis := make([]bool, n)
	q = append(q, src)
	vis[src] = true
	d := make([]int, n)
	d[src] = 0
	for len(q) > 0 {
		nn := len(q)
		for i := 0; i < nn; i++ {
			tp := q[0]
			q = q[1:]
			for _, nei := range g[tp] {
				if !vis[nei] {
					vis[nei] = true
					d[nei] = 1 + d[tp]
					q = append(q, nei)
				}
			}
		}
	}
	return d
}

func countOfPairs(n int, x int, y int) []int {
	g := make(map[int][]int)
	for i := 0; i < n-1; i++ {
		g[i] = append(g[i], i+1)
		g[i+1] = append(g[i+1], i)
	}
	g[x-1] = append(g[x-1], y-1)
	g[y-1] = append(g[y-1], x-1)

	ans := make([]int, n)

	for i := 0; i < n; i++ {
		arr := dist(i, n, g)
		//fmt.Println(arr)
		for _, v := range arr {
			if v > 0 {
				ans[v-1]++
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(countOfPairs(3, 1, 3)) // Output: [6,0,0]
	fmt.Println(countOfPairs(5, 2, 4)) // Output: [10,8,2,0,0]
	fmt.Println(countOfPairs(4, 1, 1)) // Output: [6,4,2,0]
}
