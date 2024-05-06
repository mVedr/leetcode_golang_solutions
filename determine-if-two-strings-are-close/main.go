package main

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	f1, f2 := make([]int, 26), make([]int, 26)
	for _, c := range word1 {
		f1[int(c-'a')]++
	}
	for _, c := range word2 {
		f2[int(c-'a')]++
	}
	for I := range f1 {
		if f1[I] != 0 && f2[I] == 0 {
			return false
		}
		if f1[I] == 0 && f2[I] != 0 {
			return false
		}
	}
	c := make(map[int]int)
	for _, f := range f1 {
		if f > 0 {
			c[f]++
		}
	}
	for _, f := range f2 {
		if f > 0 {
			c[f]--
		}
	}
	for _, v := range c {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {

}
