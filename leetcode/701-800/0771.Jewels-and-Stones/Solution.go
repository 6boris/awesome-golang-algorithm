package Solution

func numJewelsInStones(J string, S string) int {
	m, ans := make(map[rune]int), 0

	for _, v := range S {
		m[v]++
	}

	for _, v := range J {
		ans += m[v]
	}

	return ans
}

func numJewelsInStones2(J string, S string) int {
	ans := 0
	for _, i := range J {
		for _, j := range S {
			if i == j {
				ans++
			}
		}
	}
	return ans
}
