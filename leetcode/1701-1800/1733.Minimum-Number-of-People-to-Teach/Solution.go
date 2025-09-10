package Solution

func Solution(n int, languages [][]int, friendships [][]int) int {
	cncon := make(map[int]bool)
	for _, friendship := range friendships {
		mp := make(map[int]bool)
		conm := false
		for _, lan := range languages[friendship[0]-1] {
			mp[lan] = true
		}
		for _, lan := range languages[friendship[1]-1] {
			if mp[lan] {
				conm = true
				break
			}
		}
		if !conm {
			cncon[friendship[0]-1] = true
			cncon[friendship[1]-1] = true
		}
	}

	maxCnt := 0
	cnt := make([]int, n+1)
	for person := range cncon {
		for _, lan := range languages[person] {
			cnt[lan]++
			maxCnt = max(maxCnt, cnt[lan])
		}
	}

	return len(cncon) - maxCnt
}
