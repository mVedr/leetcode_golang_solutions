package main

func isArraySpecial(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	for I := range len(nums) - 1 {
		r1, r2 := nums[I]%2, nums[I+1]%2
		if r1 == r2 {
			return false
		}
	}
	return true
}

func main() {

}
