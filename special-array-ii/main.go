package main

func isArraySpecial(nums []int, queries [][]int) []bool {
	if len(nums) == 1 {
		a := make([]bool, len(queries))
		for I := range a {
			a[I] = true
		}
		return a
	}
	pre := []int{}
	for I := range len(nums) - 1 {
		if nums[I]%2 == nums[I+1]%2 {
			pre = append(pre, 0)
		} else {
			pre = append(pre, 1)
		}
	}
	preT := make([]int, len(pre))
	preF := make([]int, len(pre))
	if pre[0] == 0 {
		preF[0] = 1
	} else {
		preT[0] = 1
	}
	for I, v := range pre {
		if I == 0 {
			continue
		}
		preF[I] = preF[I-1]
		preT[I] = preT[I-1]
		if v == 0 {
			preF[I]++
		} else {
			preT[I]++
		}
	}
	//   fmt.Println("preF: ",preF)
	// fmt.Println("preT: ",preT)
	ans := []bool{}
	for _, q := range queries {
		i, j := q[0], q[1]
		numF := 0
		if j == 0 {
			ans = append(ans, true)
			continue
		} else {
			numF += preF[j-1]
		}
		if i > 0 {
			numF -= preF[i-1]
		}
		if numF == 0 {
			ans = append(ans, true)
		} else {
			ans = append(ans, false)
		}
	}
	return ans
}

func main() {

}
