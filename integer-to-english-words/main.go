package main

import (
	"bytes"
	"fmt"
	"strings"
)

func ok(str string) string {
	return strings.Join(strings.Fields(str), " ")
}

func numberToWords(num int) string {
	mp := make(map[int]string)
	mp[1] = "One"
	mp[2] = "Two"
	mp[3] = "Three"
	mp[4] = "Four"
	mp[5] = "Five"
	mp[6] = "Six"
	mp[7] = "Seven"
	mp[8] = "Eight"
	mp[9] = "Nine"
	mp[10] = "Ten"
	mp[11] = "Eleven"
	mp[12] = "Twelve"
	mp[13] = "Thirteen"
	mp[14] = "Fourteen"
	mp[15] = "Fifteen"
	mp[16] = "Sixteen"
	mp[17] = "Seventeen"
	mp[18] = "Eighteen"
	mp[19] = "Nineteen"
	mp[20] = "Twenty"
	mp[30] = "Thirty"
	mp[40] = "Forty"
	mp[50] = "Fifty"
	mp[60] = "Sixty"
	mp[70] = "Seventy"
	mp[80] = "Eighty"
	mp[90] = "Ninety"
	mp[100] = "Hundred"
	mp[1000] = "Thousand"
	mp[1000000] = "Million"
	mp[1000000000] = "Billion"

	buf := new(bytes.Buffer)
	for num > 0 {
		if num >= 1_000_000_000 {
			nb := num / 1_000_000_000
			num = num % 1_000_000_000
			buf.WriteString(mp[nb] + " ")
			buf.WriteString("Billion" + " ")
		} else if num >= 1_000_000 {
			nm := num / 1_000_000
			num = num % 1_000_000
			buf.WriteString(helper(nm, mp) + " ")
			buf.WriteString("Million" + " ")
		} else if num >= 1_000 {
			nt := num / 1_000
			num = num % 1_000
			buf.WriteString(helper(nt, mp) + " ")
			buf.WriteString("Thousand" + " ")
		} else {
			buf.WriteString(helper(num, mp))
			num = 0
		}
	}
	fmt.Println(ok(buf.String()))
	return strings.TrimSpace(buf.String())
}

func helper(num int, mp map[int]string) string {
	buf := new(bytes.Buffer)
	if num >= 100 {
		nb := num / 100
		num = num % 100
		buf.WriteString(mp[nb] + " ")
		buf.WriteString("Hundred" + " ")
	}
	if num >= 20 {
		nt := num / 10 * 10
		num = num % 10
		buf.WriteString(mp[nt] + " ")
	}
	if num > 0 {
		buf.WriteString(mp[num])
	}
	return ok(buf.String())
}

func main() {
	fmt.Println(numberToWords(123))
	fmt.Println(numberToWords(12345))
	fmt.Println(numberToWords(50868))
}
