package Solution

func Solution(s string) string {
	bs := []byte(s)
	index := 0
	for i := 0; i < len(s); i++ {
		if bs[i] == '*' {
			if index != 0 {
				index--
			}
			continue
		}
		bs[index] = s[i]
		index++
	}
	return string(bs[:index])
}
