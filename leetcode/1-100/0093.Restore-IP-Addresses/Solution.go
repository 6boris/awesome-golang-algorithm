package Solution

import (
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
