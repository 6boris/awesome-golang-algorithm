package Solution

import "sort"

func Solution(tokens []int, power int) int {
	sort.Ints(tokens)
	maxScore := 0
	curScore := 0
	start, end := 0, len(tokens)-1
	for start <= end {
		for ; start <= end && power >= tokens[start]; start++ {
			curScore++
			maxScore = max(maxScore, curScore)
			power -= tokens[start]
		}
		if curScore <= 0 {
			break
		}
		curScore--
		power += tokens[end]
		end--
	}

	return maxScore
}
