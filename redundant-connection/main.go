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

func findRedundantConnection(edges [][]int) []int {
	ds := &DisjointSet{
		parent: make([]int, len(edges)+1),
		size:   make([]int, len(edges)+1),
	}
	ds.Init()
	for _, E := range edges {
		if !ds.Union(E[0], E[1]) {
			return E
		}
	}

	return []int{}
}

func main() {
	fmt.Println(findRedundantConnection([][]int{{1, 2}, {1, 3}, {2, 3}}))
	fmt.Println(findRedundantConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}))
}
