package Solution

func Solution(s string, k int) string {
	size := len(s)
	oneIndies := make([]int, 0, size)
	for i := range s {
		if s[i] == '1' {
			oneIndies = append(oneIndies, i)
		}
	}
	if len(oneIndies) < k {
		return ""
	}
	var start, subLen int
	ret, minLen := s, size
	for end := k - 1; end < len(oneIndies); end++ {
		start = end + 1 - k
		subLen = oneIndies[end] - oneIndies[start] + 1
		if subLen > minLen {
			continue
		}
		subStr := s[oneIndies[start] : oneIndies[end]+1]
		if subLen == minLen {
			ret = min(ret, subStr)
		} else {
			ret = subStr
		}
		minLen = min(minLen, subLen)
	}
	return ret
}
