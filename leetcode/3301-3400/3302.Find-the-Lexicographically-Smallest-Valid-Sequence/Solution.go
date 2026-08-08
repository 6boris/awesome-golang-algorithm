package Solution

func Solution(word1 string, word2 string) []int {
	n, m := len(word1), len(word2)
	last := make([]int, m)
	for i := range last {
		last[i] = -1
	}
	j := m - 1
	for i := n - 1; i >= 0; i-- {
		if j >= 0 && word1[i] == word2[j] {
			last[j] = i
			j -= 1
		}
	}
	res := []int{}
	skip := 0
	j = 0
	for i := 0; i < n; i++ {
		if j == m {
			break
		}
		if word1[i] == word2[j] || (skip == 0 && (j == m-1 || i < last[j+1])) {
			if word1[i] != word2[j] {
				skip += 1
			}
			res = append(res, i)
			j += 1
		}
	}
	if j == m {
		return res
	}
	return []int{}
}
