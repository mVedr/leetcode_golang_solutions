package main

import "strconv"

func sumDigitDifferences(nums []int) int64 {
	arr := [10][10]int{}
	for _, n := range nums {
		ns := strconv.Itoa(n)
		for I, c := range ns {
			d := int(c - 'a')
			arr[I][d]++
		}
	}
	ans := int64(0)
	for I := range arr {
		for d := 0; d <= 9; d++ {
			dc := arr[I][d]
			for cd := 0; cd <= 9; cd++ {
				if d != cd {
					cdc := arr[I][cd]
					ans += int64(dc * cdc)

				}
			}
		}
	}
	return ans
}

func main() {

}
