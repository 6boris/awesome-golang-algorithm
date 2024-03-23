package Solution

func Solution(tiles string) int {
	count := [26]int{}
	for _, c := range tiles {
		count[c-'A']++
	}
	var dfs func() int
	dfs = func() int {
		r := 1
		for i := 0; i < 26; i++ {
			if count[i] == 0 {
				continue
			}
			count[i]--
			r += dfs()
			count[i]++
		}
		return r

	}
	return dfs() - 1
}
