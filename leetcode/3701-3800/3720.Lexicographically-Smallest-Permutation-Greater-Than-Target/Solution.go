package Solution

import "sort"

func nextBytePermutation(target string) string {
	size := len(target)
	bs := []byte(target)
	pos := size - 1
	for ; pos >= 1 && bs[pos] <= bs[pos-1]; pos-- {
	}
	if pos == 0 {
		return ""
	}
	// next permutation
	for j := size - 1; j >= pos; j-- {
		if bs[j] > bs[pos-1] {
			bs[j], bs[pos-1] = bs[pos-1], bs[j]
			break
		}
	}
	part := bs[pos:]
	sort.Slice(part, func(i, j int) bool {
		return part[i] < part[j]
	})
	return string(bs)

}
func Solution(s string, target string) string {
	count := [26]int{}
	for i := range s {
		count[s[i]-'a']++
	}
	size := len(target)
	ret := make([]byte, size)
	index, i := 0, byte('0')
	for ; index < size; index++ {
		i = target[index] - 'a'
		if count[i] == 0 {
			break
		}
		count[i]--
		ret[index] = target[index]
	}

	if index == len(target) {
		return nextBytePermutation(target)
	}

	detect := func() (string, bool) {
		tmp, pos := count, index
		j := target[pos] - 'a' + 1
		for ; j < 26; j++ {
			if tmp[j] > 0 {
				ret[pos] = 'a' + j
				tmp[j], pos = tmp[j]-1, pos+1
				break
			}
		}
		if j == 26 {
			return "", false
		}

		for i := range 26 {
			for ; tmp[i] > 0; tmp[i], pos = tmp[i]-1, pos+1 {
				ret[pos] = byte(i) + 'a'
			}
		}
		return string(ret), true
	}
	got, ok := detect()
	if ok {
		return got
	}

	index--
	for ; index >= 0; index-- {
		count[target[index]-'a']++
		got, ok = detect()
		if ok {
			return got
		}
	}
	return ""
}
