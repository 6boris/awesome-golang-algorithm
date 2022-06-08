package Solution

func Solution(words []string) int {
	ans := 0
	for idx := 0; idx < len(words)-1; idx++ {
		for next := idx + 1; next < len(words); next++ {
			vec := [2]int{0, 0}
			lidx, lnext := len(words[idx]), len(words[next])
			for b := 0; b < lidx; b++ {
				vec[0] |= 1 << int(words[idx][b]-'a')
			}
			for b := 0; b < lnext; b++ {
				vec[1] |= 1 << int(words[next][b]-'a')
			}
			if vec[0]&vec[1] == 0 && lidx*lnext > ans {
				ans = lidx * lnext
			}
		}
	}
	return ans
}
