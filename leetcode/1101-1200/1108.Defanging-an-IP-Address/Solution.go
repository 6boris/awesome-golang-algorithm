package Solution

func Solution(address string) string {
	ans := make([]byte, len(address)+6)
	index := 0
	for _, b := range address {
		if b == '.' {
			ans[index] = '['
			ans[index+1] = '.'
			ans[index+2] = ']'
			index += 3
			continue
		}
		ans[index] = byte(b)
		index++
	}
	return string(ans)
}
