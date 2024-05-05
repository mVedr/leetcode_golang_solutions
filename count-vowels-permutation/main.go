package main

var (
	MOD = 1_000_000_000 + 7
	mp  = map[byte]string{
		'a': "e",
		'e': "ai",
		'i': "aeou",
		'o': "iu",
		'u': "a",
	}
)

func solve(i int, prev int, n int, dp *[][]int) int {
	if i == n {
		return 1
	}
	if (*dp)[i][prev+1] != -1 {
		return (*dp)[i][prev+1]
	}
	ans := 0
	if i == 0 {
		for c := range mp {
			ans = (ans + solve(i+1, int(byte(c)-'a'), n, dp)) % MOD
		}
	} else {
		for _, chr := range mp[byte(prev+'a')] {
			ans = (ans + solve(i+1, int(byte(chr)-'a'), n, dp)) % MOD
		}
	}
	(*dp)[i][prev+1] = ans
	return ans
}

func Init2d[T any](n, m int, val T) [][]T {
	arr := make([][]T, n)
	for i := range n {
		arr[i] = make([]T, m)
	}
	for i := range n {
		for j := range m {
			arr[i][j] = val
		}
	}
	return arr
}

func countVowelPermutation(n int) int {
	dp := Init2d(n+1, 30, -1)
	return solve(0, 0, n, &dp)
}

func main() {

}
