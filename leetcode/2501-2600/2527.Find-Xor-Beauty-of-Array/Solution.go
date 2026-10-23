package Solution

func Solution(nums []int) int {
	n := int64(len(nums))
	var finalAns int64 = 0

	for b := 0; b < 30; b++ {
		var c1 int64 = 0
		for _, num := range nums {
			if (int64(num)>>b)&1 == 1 {
				c1++
			}
		}
		c0 := n - c1

		validPairs := n*n - c0*c0

		totalTriples := validPairs * c1

		if totalTriples%2 != 0 {
			finalAns |= (1 << b)
		}
	}

	return int(finalAns)
}
