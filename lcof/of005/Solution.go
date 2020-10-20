package Solution

func replaceSpace(s string) string {
	ans := make([]rune, len(s)*3)

	i := 0
	for _, v := range s {
		if v != ' ' {
			ans[i] = v
			i++
		} else {
			ans[i], ans[i+1], ans[i+2] = '%', '2', '0'
			i += 3
		}
	}
	return string(ans)[:i]
}
