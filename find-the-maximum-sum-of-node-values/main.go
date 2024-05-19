package main

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func maximumValueSum(nums []int, k int, edges [][]int) int64 {
	sum := int64(0)
	cnt := 0
	sac := 1_000_000_000
	for _, v := range nums {
		sum += int64(max(v^k, v))
		if v^k > v {
			cnt++
		}
		sac = min(sac, abs(v-(v^k)))
	}
	if cnt%2 == 1 {
		return sum - int64(sac)
	}
	return sum
}

func main() {

}
