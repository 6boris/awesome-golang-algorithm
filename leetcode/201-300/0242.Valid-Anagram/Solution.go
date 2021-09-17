package Solution

import "sort"

func isAnagram_1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make([]rune, 256)
	for _, v := range s {
		m[v] += 1
	}
	for _, v := range t {
		m[v] -= 1
		if m[v] < 0 {
			return false
		}
	}
	return true
}

func isAnagram_2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sl, tl := []byte(s), []byte(t)
	sort.Slice(sl, func(i, j int) bool {
		return sl[i] < sl[j]
	})
	sort.Slice(tl, func(i, j int) bool {
		return tl[i] < tl[j]
	})
	for i, elem := range sl {
		if elem != tl[i] {
			return false
		}
	}
	return true
}

//	用异或计算
func isAnagram_3(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var x, y rune

	for _, v := range s {
		x ^= v
		y += v * v
	}
	for _, v := range t {
		x ^= v
		y -= v * v
	}
	return x == 0 && y == 0
}
