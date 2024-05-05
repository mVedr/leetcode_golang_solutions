package main

func solve(nums []int) int {
	ans := 0
	sum := 0
	for _, n := range nums {
		if sum+n < 0 {
			sum = 0
		} else {
			sum += n
			ans = max(ans, sum)
		}
	}
	return ans
}

func maxAbsoluteSum(nums []int) int {
	ans := 0
	ans = max(ans, solve(nums))
	for I := range nums {
		nums[I] *= -1
	}
	ans = max(ans, solve(nums))
	return ans
}

func main() {

}
