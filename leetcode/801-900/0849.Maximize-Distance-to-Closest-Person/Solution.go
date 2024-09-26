package Solution

func Solution(seats []int) int {
	ans := 0
	pre := -1
	i := 0
	for ; i < len(seats); i++ {
		s := seats[i]
		if s == 0 {
			continue
		}
		if pre == -1 {
			ans = i
		} else {
			diff := i - pre - 1
			add := 0
			if diff&1 == 1 {
				add = 1
			}
			diff = diff/2 + add
			ans = max(ans, diff)
		}
		pre = i
	}
	if seats[i-1] != 1 {
		diff := i - pre - 1
		ans = max(ans, diff)
	}
	return ans
}
