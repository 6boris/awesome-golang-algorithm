package Solution

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	if candidates == nil || len(candidates) == 0 {
		return result
	}
	sort.Ints(candidates)
	set := []int{}
	explore(candidates, target, 0, &result, set)

	return result
}

func explore(candidates []int, target int, pos int, result *[][]int, set []int) {
	if target < 0 {
		return
	}

	if target == 0 {
		c := make([]int, len(set))
		copy(c, set)
		*result = append(*result, c)
		return
	}

	for i := pos; i < len(candidates); i++ {
		set = append(set, candidates[i])
		explore(candidates, target-candidates[i], i, result, set)
		set = set[:len(set)-1]
	}
}
