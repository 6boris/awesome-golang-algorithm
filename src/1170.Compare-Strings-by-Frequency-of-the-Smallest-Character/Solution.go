package Solution

import "sort"

func Solution(queries []string, words []string) []int {
	wordsFreq := make([]int, 0)
	for _, word := range words {
		wordsFreq = append(wordsFreq, calculateFrequency(word))
	}
	sort.Ints(wordsFreq)
	res, n := make([]int, 0), len(wordsFreq)
	for _, query := range queries {
		queryFreq := calculateFrequency(query)
		res = append(res, n-sort.Search(n, func(i int) bool {
			return wordsFreq[i] > queryFreq
		}))
	}
	return res
}

func calculateFrequency(word string) int {
	var min byte = 'z'
	c := 0
	for i := 0; i < len(word); i++ {
		if word[i] < min {
			min = word[i]
			c = 0
		}
		if word[i] == min {
			c++
		}
	}
	return c
}
