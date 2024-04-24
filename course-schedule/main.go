package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	indeg := make([]int, numCourses)
	g := make([][]int, numCourses)
	for _, p := range prerequisites {
		indeg[p[1]]++
		g[p[0]] = append(g[p[0]], p[1])
	}
	q := []int{}
	for I := range numCourses {
		if indeg[I] == 0 {
			q = append(q, I)
		}
	}

	for len(q) > 0 {
		tp := q[0]
		q = q[1:]
		for _, neigh := range g[tp] {
			if indeg[neigh] > 0 {
				if indeg[neigh] == 1 {
					q = append(q, neigh)
				}
				indeg[neigh]--
			}
		}
	}
	for _, val := range indeg {
		if val != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
}
