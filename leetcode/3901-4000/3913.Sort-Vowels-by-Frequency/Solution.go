package Solution

import "sort"

func Solution(s string) string {
	freq := make(map[byte][2]int)
	for i := range s {
		if !(s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u') {
			continue
		}
		v, ok := freq[s[i]]
		if !ok {
			v = [2]int{i, 0}
		}
		v[1]++
		freq[s[i]] = v
	}
	keys := []byte{'a', 'e', 'i', 'o', 'u'}
	sort.Slice(keys, func(i, j int) bool {
		a, b := freq[keys[i]], freq[keys[j]]
		if a[1] == b[1] {
			return a[0] < b[0]
		}
		return a[1] > b[1]
	})

	ret := []byte(s)
	index := 0

	for i := range keys {
		cnt := freq[keys[i]]
		for ; index < len(ret) && cnt[1] > 0; index++ {
			if _, ok := freq[ret[index]]; !ok {
				continue
			}
			ret[index] = keys[i]
			cnt[1]--
		}

	}
	return string(ret)

}
