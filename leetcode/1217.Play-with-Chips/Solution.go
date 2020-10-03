package Solution

func Solution(chips []int) int {
	e, o := 0, 0
	for _, val := range chips {
		if val%2 == 0 {
			e++
		} else {
			o++
		}
	}
	if e < o {
		return e
	}
	return o
}
