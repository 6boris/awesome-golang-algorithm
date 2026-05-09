package Solution

func Solution(moves string) int {
	score, both := 0, 0
	for i := range moves {
		if moves[i] == 'L' {
			score++
			continue
		}
		if moves[i] == 'R' {
			score--
			continue
		}
		both++
	}
	if score < 0 {
		score = -score
	}
	return score + both
}
