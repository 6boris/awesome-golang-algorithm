package Solution

func Solution(A []int) []int {
	even, odd := 0, 1
	res := make([]int, len(A))
	for _, val := range A {
		if val%2 == 0 {
			res[even] = val
			even += 2
		} else {
			res[odd] = val
			odd += 2
		}
	}
	return res
}
