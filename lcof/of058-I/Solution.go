package Solution

import "strings"

func reverseWords(s string) string {
	tmp := strings.Trim(s, " ")
	start, end, ans := len(tmp)-1, len(tmp)-1, ""
	for start >= 0 {
		for start >= 0 && tmp[start] != ' ' {
			start--
		}
		ans += tmp[start+1:end+1] + " "
		for start >= 0 && tmp[start] == ' ' {
			start--
		}
		end = start
	}
	return strings.Trim(ans, " ")
}
