package Solution

func isMatch1023(str, pattern string) bool {
	pi := 0
	for i := 0; i < len(str); i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			if pi == len(pattern) || str[i] != pattern[pi] {
				return false
			}
			pi++
			continue
		}

		if pi != len(pattern) && str[i] == pattern[pi] {
			pi++
		}
	}
	return pi == len(pattern)
}

func Solution(queries []string, pattern string) []bool {
	ans := make([]bool, len(queries))
	for i, s := range queries {
		ans[i] = isMatch1023(s, pattern)
	}
	return ans
}
