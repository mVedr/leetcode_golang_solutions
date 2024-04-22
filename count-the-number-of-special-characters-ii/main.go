package main

func numberOfSpecialChars(word string) int {
	mpp := map[rune]int{}
	for i, ch := range word {
		if ch >= 'a' && ch <= 'z' {
			mpp[ch] = i
		} else {
			if _, ok := mpp[ch]; !ok {
				mpp[ch] = i
			}
		}
	}
	ans := 0
	for ch := 'A'; ch <= 'Z'; ch++ {
		if I, ok := mpp[ch]; ok {
			lw := ch - 'A' + 'a'
			if J, ok1 := mpp[lw]; ok1 {
				if I > J {
					ans++
				}
			}
		}
	}
	return ans
}

func main() {

}
