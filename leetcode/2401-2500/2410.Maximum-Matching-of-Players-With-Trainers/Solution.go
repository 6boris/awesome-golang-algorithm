package Solution

import "sort"

func Solution(players []int, trainers []int) int {
	sort.Ints(players)
	sort.Ints(trainers)
	i, j := 0, 0
	ans := 0
	for i < len(players) && j < len(trainers) {
		if trainers[j] >= players[i] {
			i, j = i+1, j+1
			ans++
			continue
		}
		j++
	}
	return ans
}
