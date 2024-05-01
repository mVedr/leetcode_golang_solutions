package main

func reversePrefix(word string, ch byte) string {
	I := -1
	for i, c := range word {
		if byte(c) == ch {
			I = i
			break
		}
	}
	if I == -1 {
		return word
	}

	arr := []rune(word)
	for i, j := 0, I; i < j; i, j = i+1, j-1 {

		arr[i], arr[j] = arr[j], arr[i]
	}

	return string(arr)
}

func main() {

}
