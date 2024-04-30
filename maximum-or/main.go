package main

func maximumOr(nums []int, k int) int64 {
	bt := [64]int{}
	for _, nn := range nums {
		for I := range 64 {
			if 1<<I&nn > 0 {
				bt[I]++
			}
		}
	}
	ans := int64(0)
	for _, nn := range nums {
		OrWithoutNum := int64(0)
		temp := bt
		for I := range 64 {
			if (1<<I)&nn > 0 {
				temp[I]--
			}
		}
		for I := range temp {
			if temp[I] > 0 {
				OrWithoutNum |= (1 << I)
			}
		}
		mul := int64(nn) << k
		ans = max(ans, OrWithoutNum|mul)
	}
	return ans
}

func main() {

}
