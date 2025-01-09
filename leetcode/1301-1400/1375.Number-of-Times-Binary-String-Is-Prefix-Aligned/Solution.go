package Solution

func Solution(flips []int) int {
	ans := 0
	sum, cur := 0, 0
	for i, f := range flips {
		sum += i + 1
		cur += f
		if sum == cur {
			ans++
		}
	}
	return ans
}
