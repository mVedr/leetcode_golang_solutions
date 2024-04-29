package main

import "fmt"

func canMakeSquare(grid [][]byte) bool {

	for I := range 2 {

		for K := range 2 {
			c1 := 0
			if grid[I][K] == 'B' {
				c1++
			} else {
				c1--
			}
			if grid[I][K+1] == 'B' {
				c1++
			} else {
				c1--
			}
			if grid[I+1][K] == 'B' {
				c1++
			} else {
				c1--
			}
			if grid[I+1][K+1] == 'B' {
				c1++
			} else {
				c1--
			}
			if c1 == -2 || c1 == 2 || c1 == 4 || c1 == -4 {
				return true
			}
		}

	}
	return false
}

func main() {
	fmt.Println(fmt.Println(canMakeSquare([][]byte{{'B', 'W', 'B'}, {'B', 'W', 'W'}, {'B', 'W', 'B'}})))
}
