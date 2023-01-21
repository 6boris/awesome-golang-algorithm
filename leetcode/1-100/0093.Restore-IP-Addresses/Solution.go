package Solution

import (
	"fmt"
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	var result []string
	dfs(s, []string{}, &result)
	return result
}

func dfs(s string, temp []string, result *[]string) {
	if len(temp) == 4 && len(s) == 0 {
		*result = append(*result, strings.Join(temp, "."))
		return
	}

	if len(temp) == 4 || len(s) > 3*(4-len(temp)) || len(s) < (4-len(temp)) {
		return
	}

	for i := 1; i <= 3; i++ {
		if i > len(s) {
			continue
		}
		num, _ := strconv.Atoi(s[:i])
		if s[:i] != strconv.Itoa(num) || num > 255 {
			continue
		}

		temp = append(temp, s[:i])
		dfs(s[i:], temp, result)
		temp = temp[:len(temp)-1]
	}
}

func Solution(s string) []string {
	ans := make([]string, 0)
	var dfs func(int, int, string)
	dfs = func(start, dot int, prefix string) {
		if start >= len(s) {
			return
		}
		if dot == 0 {
			diff := len(s) - start
			if diff != 1 && s[start] == '0' {
				return
			}
			now := s[start:]

			now1, _ := strconv.Atoi(now)
			if now1 >= 0 && now1 <= 255 {
				ans = append(ans, prefix+now)
			}
			return
		}
		if s[start] == '0' {
			dfs(start+1, dot-1, prefix+"0.")
			return
		}

		for end := start; end <= start+2 && end < len(s); end++ {
			now := s[start : end+1]
			now1, _ := strconv.Atoi(s[start : end+1])
			if now1 >= 0 && now1 <= 255 {
				dfs(end+1, dot-1, prefix+fmt.Sprintf("%s.", now))
			}
		}
	}

	dfs(0, 3, "")
	return ans
}
