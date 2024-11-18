package Solution

func Solution(code []int, k int) []int {
	l := len(code)
	sum := make([]int, l)
	if k == 0 {
		return sum
	}
	sum[0] = code[0]
	for i := 1; i < l; i++ {
		sum[i] = sum[i-1] + code[i]
	}
	ans := make([]int, l)
	if k < 0 {
		k = -k
		for i := range l {
			diff := i - k
			if diff > 0 {
				start, end := diff, i-1
				ans[i] = sum[end]
				if start > 0 {
					ans[i] -= sum[start-1]
				}
			} else {
				if i > 0 {
					ans[i] += sum[i-1]
				}
				ans[i] += sum[l-1] - sum[l+diff-1]
			}
		}
	} else {
		for i := range l {
			diff := i + k
			if diff < l {
				ans[i] = sum[diff] - sum[i]
			} else {
				ans[i] += sum[l-1] - sum[i]
				ans[i] += sum[diff-l]
			}
		}

	}

	return ans
}
