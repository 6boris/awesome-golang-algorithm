package Solution

func Solution(s string) int {
	ret := 0
	for i := range s {
		ret += int('z'-s[i]+1) * (i + 1)
	}
	return ret
}
