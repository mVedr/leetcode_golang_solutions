package main

func satisfiesConditions(grid [][]int) bool {
	for I := range grid {
		for J := range grid[0] {
			if J < len(grid[0])-1 {
				if grid[I][J] == grid[I][J+1] {
					return false
				}
			}
			if I < len(grid)-1 {
				if grid[I][J] != grid[I+1][J] {
					return false
				}
			}
		}
	}
	return true
}

func main() {

}
