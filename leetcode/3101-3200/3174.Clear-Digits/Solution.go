package Solution

func Solution(s string) string {
	bs := []byte(s)
	index := -1
	for i := range len(bs) {
		if !(bs[i] >= '0' && bs[i] <= '9') {
			index++
			bs[index] = bs[i]
			continue
		}
		index--
	}
	return string(bs[:index+1])
}
