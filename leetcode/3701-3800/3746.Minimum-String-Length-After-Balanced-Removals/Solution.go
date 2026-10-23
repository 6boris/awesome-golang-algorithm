package Solution

func Solution(s string) int {
	a, b := 0, 0
	for i := range s {
		if s[i] == 'a' {
			a++
			continue
		}
		b++
	}
	left := a - b
	if left < 0 {
		left = -left
	}
	return left
}
