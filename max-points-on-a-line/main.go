package main

import "fmt"

func maxPoints(points [][]int) int {
	if len(points) == 1 {
		return 1
	}
	ans := 0
	for I := range points {
		for J := range points {
			if I == J {
				continue
			}
			if points[I][0] == points[J][0] {
				continue
			} else {
				if points[I][0] == points[J][0] {
					continue
				}
				sl := float64(points[I][1]-points[J][1]) / float64(points[I][0]-points[J][0])
				temp := 0
				for K := range points {
					if I != K && J != K && points[I][0] != points[K][0] {
						sln := float64(points[I][1]-points[K][1]) / float64(points[I][0]-points[K][0])
						if sl == sln {
							temp++
						}
					}
				}
				ans = max(temp+2, ans)
			}
		}
	}
	for _, v1 := range points {
		temp := 0
		XX := v1[0]
		for _, v2 := range points {
			if v2[0] == XX {
				temp++
			}
		}
		ans = max(ans, temp)
	}
	return ans
}

func main() {
	fmt.Println(maxPoints([][]int{{1, 1}, {2, 2}, {3, 3}}))
	fmt.Println(maxPoints([][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}))
}
