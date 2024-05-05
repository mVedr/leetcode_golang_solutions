package main

func ok(bt []int) bool {

	for _, v := range bt {
		if v > 1 {
			return false
		}
	}

	return true
}

func longestNiceSubarray(nums []int) int {
	ans := 0
	bt := make([]int, 32)
	I, J := 0, 0

	for J < len(nums) {
		for k := range 32 {
			if (1<<k)&nums[J] > 0 {
				bt[k]++
			}
		}
		for I != J && !ok(bt) {
			for k := range 32 {
				if (1<<k)&nums[I] > 0 {
					bt[k]--
				}
			}
			I++
		}
		J++
	}
	if ok(bt) {
		ans = max(ans, J-I+1)
	}
	return ans
}

func main() {

}
