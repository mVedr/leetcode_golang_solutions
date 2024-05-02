package main

func trap(height []int) int {
	ans := 0
	pre, suf := make([]int, len(height)), make([]int, len(height))
	for I := range height {
		if I == 0 {
			pre[I] = height[I]
		}
		pre[I] = max(pre[I-1], height[I])
	}
	for J := len(height) - 1; J >= 0; J-- {
		if J == len(height)-1 {
			suf[J] = height[J]
			continue
		}
		suf[J] = max(suf[J+1], height[J])
	}
	for K := 1; K < len(pre)-1; K++ {
		ans += min(pre[K-1], suf[K-1]) - height[K]
	}
	return ans
}

func main() {

}
