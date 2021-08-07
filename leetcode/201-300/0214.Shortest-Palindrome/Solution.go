package Solution

func Solution(s string) string {
	length := len(s)
	reverse := make([]byte, length)
	for idx := length - 1; idx >= 0; idx-- {
		reverse[length-idx-1] = s[idx]
	}

	str := string(reverse)
	idx := length
	for ; idx > 0; idx-- {
		if s[:idx] == str[length-idx:] {
			break
		}
	}
	if idx == length {
		return s
	}

	return str[:length-idx] + s
}
