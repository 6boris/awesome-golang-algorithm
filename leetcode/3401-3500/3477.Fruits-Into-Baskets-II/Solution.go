package Solution

func Solution(fruits, baskets []int) int {
	left := len(fruits)
	for _, f := range fruits {
		for i := range baskets {
			if baskets[i] >= f {
				left--
				baskets[i] = -1
				break
			}
		}
	}
	return left
}
