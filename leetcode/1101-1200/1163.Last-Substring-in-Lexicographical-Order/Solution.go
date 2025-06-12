package Solution

func Solution(s string) string {
	indies := [26][]int{}
	for i, b := range s {
		indies[b-'a'] = append(indies[b-'a'], i)
	}
	i := 25
	for ; i >= 0 && len(indies[i]) == 0; i-- {
	}
	ans := ""
	for _, index := range indies[i] {
		ans = max(ans, s[index:])
	}
	return ans
}
