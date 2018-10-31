package Solution

func longestCommonPrefix(strs []string) string {
	var result string

	maxIndex := 0
	for i, str := range strs {
		strLen := len(str)
		if i == 0 {
			maxIndex = strLen - 1
			result = str
			continue
		}

		if strLen-1 < maxIndex {
			maxIndex = strLen - 1
			result = result[:strLen]
		}

		for j := 0; j <= maxIndex && j < strLen; j++ {
			if str[j] != result[j] {
				maxIndex = j - 1
				result = str[:j]
			}
		}
	}

	return result
}
