package Solution

func Solution(s string) int {
	countOne := 0
	ans := 0
	i := 0
	for i < len(s) {
		if s[i] == '0' {
			for i+1 < len(s) && s[i+1] == '0' {
				i++
			}
			ans += countOne
		} else {
			countOne++
		}
		i++
	}
	return ans
}
