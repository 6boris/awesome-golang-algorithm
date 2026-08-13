package Solution

import "sort"

func Solution(word string) int {
	cnt := [26]int{}
	indies := make([]int, 26)
	for i := range 26 {
		indies[i] = i
	}
	for i := range word {
		cnt[word[i]-'a']++
	}
	sort.Slice(indies, func(i, j int) bool {
		ii, jj := indies[i], indies[j]
		return cnt[ii] > cnt[jj]
	})

	repeat := 0
	var ret int
	for i := 0; i < 26 && cnt[indies[i]] > 0; i, repeat = i+1, repeat+1 {
		ret += cnt[indies[i]] * (repeat/8 + 1)
	}
	return ret
}
