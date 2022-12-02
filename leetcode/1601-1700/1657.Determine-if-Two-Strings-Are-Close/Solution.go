package Solution

import "sort"

func Solution(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	bucket1, bucket2 := make([]int, 26), make([]int, 26)
	for idx := 0; idx < len(word1); idx++ {
		bucket1[word1[idx]-'a']++
		bucket2[word2[idx]-'a']++
	}

	for i := 0; i < 26; i++ {
		if bucket2[i] != 0 && bucket1[i] == 0 || bucket1[i] != 0 && bucket2[i] == 0 {
			return false
		}
	}
	sort.Ints(bucket1)
	sort.Ints(bucket2)

	for i := 0; i < 26; i++ {
		if bucket1[i] != bucket2[i] {
			return false
		}
	}
	return true
}
