package Solution

func Solution(s1 string, s2 string) int {
	// 寻找diff的情况
	if len(s1) != len(s2) {
		return -1
	}
	zero, one := 0, 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			if s1[i] == 'x' {
				// xy
				zero++
			} else {
				// yx
				one++
			}
		}
	}
	if (one+zero)&1 == 1 {
		return -1
	}
	ans := one / 2
	ans += zero / 2
	if one&1 == 1 {
		ans += 2
	}
	return ans
}
