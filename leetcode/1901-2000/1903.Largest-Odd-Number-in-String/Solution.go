package Solution

func Solution(num string) string {
	idx := len(num) - 1
	for ; idx >= 0; idx-- {
		if num[idx]&1 == 1 {
			break
		}
	}
	return num[:idx+1]
}
