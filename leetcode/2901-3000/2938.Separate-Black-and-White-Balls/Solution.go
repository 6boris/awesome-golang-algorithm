package Solution

func Solution(s string) int64 {
	index := -1
	dis := int64(0)
	for i := range s {
		if s[i] == '0' {
			index++
			dis += int64(i - index)
		}
	}
	return dis
}
