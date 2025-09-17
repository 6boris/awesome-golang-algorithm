package Solution

func gcd2197(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm2197(a, b, gcd int) int {
	return a / gcd * b
}

func Solution(nums []int) []int {
	l := len(nums)
	stack := make([]int, l)
	stack[0] = nums[0]
	index := 0
	for i := 1; i < len(nums); i++ {
		g := gcd2197(stack[index], nums[i])
		// 如果可以一只gcd下去，那就一只gcd2197
		cmp := nums[i]
		for ; index >= 0; index-- {
			g = gcd2197(stack[index], cmp)
			if g == 1 {
				break
			}
			cmp = lcm2197(stack[index], cmp, g)
		}
		index++
		stack[index] = cmp
	}
	return stack[:index+1]
}
