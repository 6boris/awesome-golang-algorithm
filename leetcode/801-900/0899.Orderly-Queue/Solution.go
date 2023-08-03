package Solution

import "sort"

func Solution(s string, k int) string {
	l := len(s)
	bs := []byte(s)
	// abc k = 2
	// bcad k = 2
	// cadb
	// adbc
	// abcd

	// dcba k = 2
	// cbad
	// badc
	// bdca
	// dcab

	if k != 1 {
		sort.Slice(bs, func(i, j int) bool {
			return bs[i] < bs[j]
		})
		return string(bs)
	}
	ans := s
	cache := make(map[string]struct{})
	cache[s] = struct{}{}
	for {
		maxIndex := 0
		for i := 1; i < k; i++ {
			if bs[i] >= bs[maxIndex] {
				maxIndex = i
			}
		}
		sourceByte := bs[maxIndex]
		for next := maxIndex + 1; next < l; next++ {
			bs[next-1] = bs[next]
		}
		bs[l-1] = sourceByte
		r := string(bs)
		if _, ok := cache[r]; ok {
			break
		}
		cache[r] = struct{}{}
		if r < ans {
			ans = r
		}
	}
	return ans
}
