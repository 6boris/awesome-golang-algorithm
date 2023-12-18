package Solution

func Solution(num string) string {
	ans := ""
	if len(num) < 3 {
		return ans
	}
	a, b, c := 0, 1, 2
	if num[a] == num[b] && num[b] == num[c] {
		ans = num[a : c+1]
	}

	for idx := 3; idx < len(num); idx++ {
		a, b, c = b, c, idx
		if num[a] == num[b] && num[b] == num[c] {
			if ans == "" || ans < num[a:c+1] {
				ans = num[a : c+1]
			}
		}

	}
	return ans
}
