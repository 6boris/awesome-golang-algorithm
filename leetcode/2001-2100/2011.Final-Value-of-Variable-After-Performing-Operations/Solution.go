package Solution

func Solution(operations []string) int {

	ans := 0
	for _, opt := range operations {
		if opt == "++X" || opt == "X++" {
			ans++
			continue
		}
		ans--
	}
	return ans
}
