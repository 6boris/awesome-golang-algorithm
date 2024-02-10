package Solution

import (
	"sort"
)

func Solution(votes []string) string {
	count := [26][]int{}
	l := len(votes[0])
	sorter := []byte(votes[0])
	for i := 0; i < 26; i++ {
		count[i] = make([]int, l)
	}
	for _, vote := range votes {
		for i, c := range vote {
			count[c-'A'][i]++
		}
	}

	sort.Slice(sorter, func(i, j int) bool {
		ci := sorter[i]
		cj := sorter[j]
		for k := 0; k < l; k++ {
			if count[ci-'A'][k] < count[cj-'A'][k] {
				return false
			}
			if count[ci-'A'][k] > count[cj-'A'][k] {
				return true
			}
		}
		return ci < cj
	})
	return string(sorter)
}
