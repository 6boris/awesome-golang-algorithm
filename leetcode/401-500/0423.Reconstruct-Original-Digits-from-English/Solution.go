package Solution

import "strings"

func removeChars(a *[26]int, str string, count int) {
	for _, b := range str {
		a[b-'a'] -= count
	}
}

func Solution(s string) string {
	chars := [26]int{}
	for _, b := range s {
		chars[b-'a']++
	}
	buf := strings.Builder{}
	ans := [10]int{}
	if chars['z'-'a'] > 0 {
		ans[0] = chars[25]
		removeChars(&chars, "zero", ans[0])
	}
	if chars['w'-'a'] > 0 {
		ans[2] = chars['w'-'a']
		removeChars(&chars, "two", ans[2])
	}
	if chars['x'-'a'] > 0 {
		ans[6] = chars['x'-'a']
		removeChars(&chars, "six", ans[6])
	}
	if chars['u'-'a'] > 0 {
		ans[4] = chars['u'-'a']
		removeChars(&chars, "four", ans[4])
	}
	if chars['f'-'a'] > 0 {
		ans[5] = chars['f'-'a']
		removeChars(&chars, "five", ans[5])
	}

	if chars['v'-'a'] > 0 {
		ans[7] = chars['v'-'a']
		removeChars(&chars, "seven", ans[7])
	}
	if chars['o'-'a'] > 0 {
		ans[1] = chars['o'-'a']
		removeChars(&chars, "one", ans[1])
	}
	if chars['n'-'a'] > 0 {
		ans[9] = chars['n'-'a'] / 2
		removeChars(&chars, "nine", ans[9])
	}

	if chars['r'-'a'] > 0 {
		ans[3] = chars['r'-'a']
		removeChars(&chars, "three", ans[3])
	}
	if chars['e'-'a'] > 0 {
		ans[8] = chars['e'-'a']
	}
	for i := 0; i < 10; i++ {
		for ; ans[i] > 0; ans[i]-- {
			buf.WriteByte(byte(i + 48))
		}
	}
	return buf.String()
}
