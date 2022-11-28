package Solution

func Solution(matches [][]int) [][]int {
	ans := make([][]int, 2)
	ans[0] = make([]int, 0)
	ans[1] = make([]int, 0)
	winners, losers := make([]int, 100001), make([]int, 100001)
	for _, match := range matches {
		winners[match[0]]++
		losers[match[1]]++
	}

	for win := 0; win < 100001; win++ {
		if losers[win] == 1 {
			ans[1] = append(ans[1], win)
		}
		if winners[win] > 0 && losers[win] == 0 {
			ans[0] = append(ans[0], win)
		}
	}
	return ans
}
