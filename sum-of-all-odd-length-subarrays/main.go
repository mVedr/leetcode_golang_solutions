package main

func sumOddLengthSubarrays(arr []int) int {
	ans := 0
	for I := range arr {
		ans += ((len(arr)-I)*(I+1) + 1) / 2 * arr[I]
	}
	return ans
}

func main() {

}
