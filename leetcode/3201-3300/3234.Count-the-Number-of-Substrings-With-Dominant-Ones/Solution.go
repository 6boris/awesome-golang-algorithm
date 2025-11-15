package Solution

func Solution(s string) int {
	n := len(s)
	pre := make([]int, n+1)
	pre[0] = -1
	for i := 0; i < n; i++ {
		if i == 0 || (i > 0 && s[i-1] == '0') {
			pre[i+1] = i
		} else {
			pre[i+1] = pre[i]
		}
	}
	res := 0
	for i := 1; i <= n; i++ {
		cnt0 := 0
		if s[i-1] == '0' {
			cnt0 = 1
		}
		j := i
		for j > 0 && cnt0*cnt0 <= n {
			cnt1 := (i - pre[j]) - cnt0
			if cnt0*cnt0 <= cnt1 {
				add := j - pre[j]
				if cnt1-cnt0*cnt0+1 < add {
					add = cnt1 - cnt0*cnt0 + 1
				}
				res += add
			}
			j = pre[j]
			cnt0++
		}
	}
	return res
}
