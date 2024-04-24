package main

import "fmt"

type DisjointSet struct {
	parent []int
	size   []int
}

func (ds *DisjointSet) Union(u, v int) bool {
	ulp_u, ulp_v := ds.FindUPar(u), ds.FindUPar(v)
	if ulp_u == ulp_v {
		return false
	}
	if (*ds).size[ulp_v] > (*ds).size[ulp_u] {
		(*ds).size[ulp_v] += (*ds).size[ulp_u]
		ds.parent[ulp_u] = ulp_v
	} else {
		(*ds).size[ulp_u] += (*ds).size[ulp_v]
		ds.parent[ulp_v] = ulp_u
	}
	return true
}

func (ds *DisjointSet) Init() {
	for i := range (*ds).parent {
		(*ds).parent[i] = i
		(*ds).size[i] = 1
	}
}

func (ds *DisjointSet) FindUPar(u int) int {
	for (*ds).parent[u] != u {
		u = (*ds).parent[u]
	}
	return u
}

func makeConnected(n int, connections [][]int) int {
	ds := &DisjointSet{
		parent: make([]int, n),
		size:   make([]int, n),
	}
	ds.Init()
	rems := 0
	for _, c := range connections {
		if ds.Union(c[0], c[1]) {

		} else {
			rems++
		}
	}
	comps := 0
	for I := range ds.parent {
		if (*ds).parent[I] == I {
			comps++
		}
	}
	if rems >= comps-1 {
		return comps - 1
	}
	return -1
}

func main() {
	fmt.Println(makeConnected(4, [][]int{{0, 1}, {0, 2}, {1, 2}}))
	fmt.Println(makeConnected(6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}}))
	fmt.Println(makeConnected(6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}}))
}
