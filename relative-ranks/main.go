package main

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	ans := []string{}
	temp := score
	sort.Slice(temp, func(i, j int) bool {
		return temp[i] < temp[j]
	})
	mp := make(map[int]int)
	for I := range temp {
		mp[temp[I]] = I + 1
	}
	for _, v := range score {
		if v == 1 {
			ans = append(ans, "Gold Medal")
		} else if v == 2 {
			ans = append(ans, "Silver Medal")
		} else if v == 3 {
			ans = append(ans, "Bronze Medal")
		} else {
			ans = append(ans, strconv.FormatInt(int64(v), 10))
		}
	}
	return ans
}

func main() {

}
