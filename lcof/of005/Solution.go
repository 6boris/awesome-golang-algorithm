package Solution

func replaceSpace(s string) string {
	result := make([]rune, len(s)*3)

	i := 0
	for _, v := range s {
		if v != ' ' {
			result[i] = v
			i++
		} else {
			result[i] = '%'
			result[i+1] = '2'
			result[i+2] = '0'
			i += 3
		}
	}
	return string(result)[:i]
}
