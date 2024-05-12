package main

import (
	"fmt"
)

func solve(i int, buf []byte, dp map[int]int) int {
	if i >= len(buf) {
		return 0
	}
	if v, ok := dp[i]; ok {
		return v
	}
	ans := 1001
	cnt := make([]int, 26)
	for j := i; j < len(buf); j++ {
		cnt[buf[j]-'a']++
		maxCnt := 0
		minCnt := 1001
		for k := 0; k < 26; k++ {
			if cnt[k] > 0 {
				maxCnt = max(maxCnt, cnt[k])
				minCnt = min(minCnt, cnt[k])
			}
		}
		if maxCnt == minCnt {
			ans = min(ans, 1+solve(j+1, buf, dp))
		}
	}
	dp[i] = ans
	return ans
}

func minimumSubstringsInPartition(s string) int {
	dp := make(map[int]int)
	return solve(0, []byte(s), dp)
}

func main() {
	fmt.Println(minimumSubstringsInPartition("fabccddg"))
	fmt.Println(minimumSubstringsInPartition("abababaccddb"))
}
