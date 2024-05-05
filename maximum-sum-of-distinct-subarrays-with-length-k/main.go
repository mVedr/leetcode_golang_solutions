package main

func maximumSubarraySum(nums []int, k int) int64 {
	ans := int64(0)
	cnt := make(map[int]int)
	I, J := 0, 0
	sum := int64(0)
	for J < len(nums) {
		if J-I+1 < k {
			sum += int64(nums[J])
			cnt[nums[J]]++
			J++
		}
		if J-I+1 == k {
			if len(cnt) == k {
				ans = max(ans, sum)
			}
			cnt[nums[I]]--
			if cnt[nums[I]] == 0 {
				delete(cnt, nums[I])
			}
			I--
		}
	}
	return ans
}

func main() {

}
