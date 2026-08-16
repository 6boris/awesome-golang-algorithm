package Solution

func Solution(stones []int) bool {
	score := [3]int{}
	for i := range stones {
		score[stones[i]%3]++
	}
	if score[0]%2 == 0 {
		return score[1] >= 1 && score[2] >= 1
	}
	return score[1]-score[2] > 2 || score[2]-score[1] > 2
}
