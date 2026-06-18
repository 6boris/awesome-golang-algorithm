package Solution

func Solution(hour int, minutes int) float64 {
	// 分针走的度数
	mi := (float64(minutes) / 60.0) * 360.0

	h := (hour % 12) * 30
	ho := float64(h) + (float64(minutes)/60.0)*30.0

	diff := mi - ho
	if diff < 0 {
		diff = -diff
	}
	return min(diff, 360.0-diff)
}
