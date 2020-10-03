package Solution

func strStr(haystack string, needle string) int {
	//	检查数据
	if len(haystack) < len(needle) {
		return -1
	}
	if len(needle) == 0 {
		return 0
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		j := 0
		for ; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}

		if j == len(needle) {
			return i
		}
	}
	return -1
}
