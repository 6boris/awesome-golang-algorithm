package Solution

import "reflect"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap1 := make([]uint8, 256)
	sMap2 := make([]uint8, 256)

	for i := 0; i < len(s); i++ {
		sMap1[s[i]] += 1
		sMap2[t[i]] += 1
	}
	if !reflect.DeepEqual(sMap1, sMap2) {
		return false
	}
	return true
}

//	先排序再比较
func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s, t = InsertSort(s), InsertSort(t)

	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			return false
		}
	}
	return true
}

// 插入排序(Insert Sort)
func InsertSort(arr string) string {
	strs := []byte(arr)

	i, j := 0, 0
	for i = 1; i < len(strs); i++ {
		tmp := strs[i]
		for j = i; j > 0 && strs[j-1] > tmp; j-- {
			strs[j] = strs[j-1]
		}
		strs[j] = tmp
	}

	return string(strs)
}

//	用异或计算
func isAnagram2(s string, t string) bool {
	var xor, squaresum1, squaresum2 rune
	for _, char := range s {
		xor ^= char
		squaresum1 += char * char
	}
	for _, char := range t {
		xor ^= char
		squaresum2 += char * char
	}
	return xor == 0 && squaresum1 == squaresum2
}

//	用Map
func isAnagram3(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap1 := make(map[uint8]int)
	sMap2 := make(map[uint8]int)

	for i := 0; i < len(s); i++ {
		sMap1[s[i]] += 1
		sMap2[t[i]] += 1
	}
	if !reflect.DeepEqual(sMap1, sMap2) {
		return false
	}
	return true
}
