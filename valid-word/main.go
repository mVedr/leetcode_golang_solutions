package main

func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}
	a := false
	b := false

	for _, ch := range word {
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' || ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' {
			a = true
		} else if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
			b = true
		} else if ch == '@' || ch == '#' || ch == '$' {
			return false
		}
	}
	return a && b
}

func main() {

}
