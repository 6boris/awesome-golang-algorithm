package Solution

import "sort"

func Solution(words []string) string {
	sort.Slice(words, func(i, j int) bool {
		li, lj := len(words[i]), len(words[j])
		if li == lj {
			return words[i] < words[j]
		}
		return li < lj
	})
	ok := make(map[string]struct{})
	ans := ""
	ok[ans] = struct{}{}
	for _, w := range words {
		lp := len(w)
		pre := w[:lp-1]
		if _, nice := ok[pre]; nice {
			ok[w] = struct{}{}
			if lp > len(ans) || (lp == len(ans) && w < ans) {
				ans = w
			}
		}
	}

	return ans
}
