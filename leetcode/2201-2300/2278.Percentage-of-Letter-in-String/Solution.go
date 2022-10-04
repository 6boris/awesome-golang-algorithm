package Solution

func Solution(s string, letter byte) int {
	count := 0
	for _, b := range s {
		if byte(b) == letter {
			count++
		}
	}

	return count * 100 / len(s)
}
