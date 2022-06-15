package Solution

import (
	"sort"
)

func Solution(words []string) int {
	length := len(words)
	dp := make([]int, length)
	for idx := 0; idx < length; idx++ {
		dp[idx] = 1
	}

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	ans := 1
	for idx := 1; idx < length; idx++ {
		for in := 0; in < idx; in++ {
			if isPredecessor(words[in], words[idx]) && dp[in]+1 > dp[idx] {
				dp[idx] = dp[in] + 1
			}
		}
		if dp[idx] > ans {
			ans = dp[idx]
		}
	}
	return ans
}

func isPredecessor(wordA, wordB string) bool {
	if len(wordB)-len(wordA) == 1 {
		ia, ib := 0, 0
		for ia < len(wordA) && ib < len(wordB) {
			if wordA[ia] == wordB[ib] {
				ia++
			}
			ib++
		}
		return ia == len(wordA)
	}
	return false
}
