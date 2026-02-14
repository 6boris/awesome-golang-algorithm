package Solution

func Solution(s string, t string) int {
	var (
		ret, start, end int

		sub string
	)
	diffOneChar := func(a, b string) bool {
		cnt := 0
		for i := range a {
			if a[i] != b[i] {
				cnt++
			}
			if cnt > 1 {
				return false
			}
		}
		return cnt == 1
	}

	ls, lt := len(s), len(t)

	for l := 1; l <= min(ls, lt); l++ {
		cache := make(map[string]int)
		subCache := make(map[string]int)
		for start := 0; start <= lt-l; start++ {
			end = start + l
			subCache[t[start:end]]++
		}
		// 1, 2, 3, 4
		for start = 0; start <= ls-l; start++ {
			end = start + l
			sub = s[start:end]
			if _, ok := cache[sub]; !ok {
				for ts, cnt := range subCache {
					if diffOneChar(sub, ts) {
						cache[sub] += cnt
					}
				}
			}

			ret += cache[sub]
		}
	}
	return ret
}
