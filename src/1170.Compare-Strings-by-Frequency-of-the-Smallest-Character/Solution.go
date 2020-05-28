package Solution

import "sort"

func Solution(queries []string, words []string) []int {
	wordsFreq := make([]int, 0)
	for _, word := range words {
		wordsFreq = append(wordsFreq, calculateFrequency(word))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(wordsFreq)))
	res := make([]int, 0)
	for _, query := range queries {
		queryFreq, c := calculateFrequency(query), 0
		for _, val := range wordsFreq {
			if queryFreq < val {
				c++
			} else {
				break
			}
		}
		res = append(res, c)
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
