package Solution

func Solution(s string) string {
	length := len(s)
	if length == 0 {
		return ""
	}
	ans := make([]byte, length)

	idx, walker := -1, 0
	for ; walker < length; walker++ {
		if idx == -1 || !byteEqual(ans[idx], s[walker]) {
			idx++
			ans[idx] = s[walker]
			continue
		}
		idx--
	}
	return string(ans[:idx+1])
}

func byteEqual(a, b byte) bool {
	diff := int(a) - int(b)
	return diff == 32 || diff == -32
}
