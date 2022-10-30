package Solution

func Solution(text string) int {
	target := "balloon"
	can := 10001

	bucket := make([]int, 26)
	for _, b := range text {
		bucket[b-'a']++
	}

	for _, b := range target {
		n := bucket[b-'a']
		if b == 'l' || b == 'o' {
			n /= 2
		}
		if n < can {
			can = n
		}
	}
	return can
}
