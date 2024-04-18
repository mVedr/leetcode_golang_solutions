package main

import "fmt"

func solve(mat [][]int) []int {
	ans := []int{}
	n, m := len(mat), len(mat[0])
	U, D, L, R := 0, n-1, 0, m-1

	for U <= D && L <= R {
		for i := L; i <= R; i++ {
			ans = append(ans, mat[U][i])
		}
		U++
		for i := U; i <= D; i++ {
			ans = append(ans, mat[i][R])
		}
		R--
		if U <= D {
			for i := R; i >= L; i-- {
				ans = append(ans, mat[D][i])
			}
			D--
		}
		if L <= R {
			for i := D; i >= U; i-- {
				ans = append(ans, mat[i][L])
			}
			L++
		}
	}

	return ans
}

func main() {
	mat1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	mat2 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	fmt.Println(solve(mat1))
	fmt.Println(solve(mat2))

}
