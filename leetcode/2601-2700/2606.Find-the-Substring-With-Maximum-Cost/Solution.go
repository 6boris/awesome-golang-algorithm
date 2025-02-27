package Solution

func Solution(s string, chars string, vals []int) int {
	value := [26]int{}
	for i := range 26 {
		value[i] = i + 1
	}
	for i, c := range chars {
		value[c-'a'] = vals[i]
	}
	vv := make([]int, len(s))
	for i := range s {
		vv[i] = value[s[i]-'a']
	}

	ans, minSum, sum := 0, 0, 0
	for i := 0; i < len(vv); i++ {
		ans = max(ans, vv[i])
		sum += vv[i]
		ans = max(ans, sum-minSum)
		if sum < minSum {
			minSum = sum
		}
	}
	return ans
}
