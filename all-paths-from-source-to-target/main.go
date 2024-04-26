package main

import "fmt"

func dfs(curr int, g [][]int, temp []int, ans *[][]int) {
	if curr == len(g)-1 {
		path := make([]int, len(temp)+1)
		temp = append(temp, curr)
		copy(path, temp)
		*ans = append(*ans, path)
		return
	}
	for _, neigh := range g[curr] {
		temp = append(temp, curr)
		dfs(neigh, g, temp, ans)
		temp = temp[:len(temp)-1]
	}
}

func allPathsSourceTarget(g [][]int) [][]int {
	ans := [][]int{}
	dfs(0, g, []int{}, &ans)
	return ans
}

func main() {
	fmt.Println(allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, {}}))
	fmt.Println(allPathsSourceTarget([][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}))
}
