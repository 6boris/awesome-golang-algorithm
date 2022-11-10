package Solution

func Solution(s string) string {
	ans := make([]byte, len(s))
	idx := -1
	for walker := 0; walker < len(s); walker++ {
		if idx == -1 || s[walker] != ans[idx] {
			idx++
			ans[idx] = s[walker]
			continue
		}
		idx--
	}
	return string(ans[:idx+1])
}
