package main

import "fmt"

func minMutation(start string, end string, bank []string) int {
	dict := make(map[string]bool)
	for _, word := range bank {
		dict[word] = true
	}

	if !dict[end] {
		return -1
	}

	str := []byte{'A', 'C', 'G', 'T'}

	vis := make(map[string]bool)
	ans := 0

	q := []string{start}
	vis[start] = true

	for len(q) > 0 {
		n := len(q)
		for n > 0 {
			curr := q[0]
			q = q[1:]
			n--

			if curr == end {
				return ans
			}

			for i := 0; i < len(curr); i++ {
				orig := curr[i]

				for _, ch := range str {
					if ch != orig {
						temp := []byte(curr)
						temp[i] = ch
						next := string(temp)

						if dict[next] && !vis[next] {
							q = append(q, next)
							vis[next] = true
						}
					}
				}
			}
		}
		ans++
	}
	return -1
}

func main() {
	fmt.Println(minMutation("AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}))
	fmt.Println(minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}))
}
