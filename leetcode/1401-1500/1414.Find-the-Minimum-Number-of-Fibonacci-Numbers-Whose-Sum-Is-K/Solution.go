package Solution

func Solution(k int) int {
	fb := []int{1, 1}
	for i := 2; i < 44 && fb[i-1]+fb[i-2] <= k; i++ {
		fb = append(fb, fb[i-1]+fb[i-2])
	}
	ret := 0
	// 5
	for index := len(fb) - 1; index >= 0 && k > 0; index-- {
		if k < fb[index] {
			continue
		}
		ret += k / fb[index]
		k %= fb[index]
	}
	return ret
}
