package Solution

func Solution(tomatoSlices int, cheeseSlices int) []int {
	// 4x + 2y = tomatoSlices
	// 2x + 2y = 2*cheeseSlices
	jumbo := tomatoSlices - 2*cheeseSlices
	if jumbo < 0 || jumbo&1 == 1 {
		return []int{}
	}
	a := jumbo / 2
	b := cheeseSlices - a
	if b < 0 {
		return []int{}
	}
	return []int{a, b}
}
