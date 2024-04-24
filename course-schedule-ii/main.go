package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	indeg := make([]int, numCourses)
	g := make([][]int, numCourses)
	for _, p := range prerequisites {
		indeg[p[0]]++
		g[p[1]] = append(g[p[1]], p[0])
	}
	q := []int{}
	ans := []int{}
	for I := range numCourses {
		if indeg[I] == 0 {
			q = append(q, I)
			ans = append(ans, I)
		}
	}

	for len(q) > 0 {
		n := len(q)
		for range n {
			tp := q[0]
			q = q[1:]
			for _, neigh := range g[tp] {
				if indeg[neigh] > 0 {
					if indeg[neigh] == 1 {
						q = append(q, neigh)
						ans = append(ans, neigh)
					}
					indeg[neigh]--
				}
			}
		}

	}
	for _, val := range indeg {
		if val != 0 {
			return []int{}
		}
	}
	return ans
}

func main() {

}
