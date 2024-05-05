package main

func solve(nums []int, k int, c int) int {
	ans := 0
	I, J := 0, 0
	for J < len(nums) {
		if nums[J] != c {
			k--
		}
		if k < 0 {
			for k < 0 {
				if nums[I] != c {
					k++
				}
				I++
			}
		} else {
			ans = max(ans, J-I+1)
		}
		J++
	}
	return ans
}

func maxConsecutiveAnswers(answerKey string, k int) int {
	nums := []int{}
	for _, ch := range answerKey {
		if ch == 'T' {
			nums = append(nums, 1)
		} else {
			nums = append(nums, 0)
		}
	}
	return max(solve(nums, k, 0), solve(nums, k, 1))
}

func main() {

}
