package Solution

func Solution(s string) string {
	r := make([]byte, len(s))
	lastOneIndex := 0
	cur := 0
	for i, b := range s {
		r[i] = '0'
		if b == '1' {
			lastOneIndex = cur
			r[cur] = '1'
			cur++
		}
	}
	r[lastOneIndex] = '0'
	r[len(r)-1] = '1'
	return string(r)
}
