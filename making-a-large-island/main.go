package main

import "fmt"

var (
	dx = []int{1, -1, 0, 0}
	dy = []int{0, 0, -1, 1}
)

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

func getMaxArea(i, j int, grid [][]int, d *DisjointSet) int {
	vis := make(map[int]bool)
	n, m := len(grid), len(grid[0])
	maxArea := 1

	for k := range 4 {
		cx, cy := i+dx[k], j+dy[k]
		if cx >= 0 && cx < n && cy >= 0 && cy < m && grid[cx][cy] == 1 {
			neigh := d.FindUPar(cx*n + cy)
			if !vis[neigh] {
				maxArea += d.size[neigh]
				vis[neigh] = true
			}
		}
	}

	return maxArea
}

func largestIsland(grid [][]int) int {
	ds := &DisjointSet{
		parent: make([]int, len(grid)*len(grid[0])),
		size:   make([]int, len(grid)*len(grid[0])),
	}
	ds.Init()
	ans := 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 {
				for k := range 4 {
					cx, cy := i+dx[k], j+dy[k]
					if cx >= 0 && cx < len(grid) && cy >= 0 && cy < len(grid[0]) && grid[cx][cy] == 1 {
						ds.Union(i*len(grid)+j, cx*len(grid)+cy)
					}
				}
			}
		}
	}

	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 0 {
				ans = max(ans, getMaxArea(i, j, grid, ds))
			}
		}
	}

	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 {
				c := ds.FindUPar(i*len(grid) + j)
				ans = max(ans, ds.size[c])
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(largestIsland([][]int{{1, 0}, {0, 1}}))
	fmt.Println(largestIsland([][]int{{1, 1}, {1, 0}}))
	fmt.Println(largestIsland([][]int{{1, 1}, {1, 1}}))
}
