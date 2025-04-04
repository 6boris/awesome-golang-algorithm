package Solution

func Solution(s string) int {
	cnt := 0
	start := -1
	end := 0
	for ; end < len(s); end++ {
		if s[end] == ' ' {
			if start == -1 {
				continue
			}
			cnt++
			start = -1
			continue
		}
		if start == -1 {
			start = end
		}
	}
	if start != -1 {
		cnt++
	}
	return cnt
}
