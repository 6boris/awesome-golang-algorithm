package Solution

func Solution(num uint32) int {
	ans := 0
	for num > 0 {
		ans++
		num = num & (num - 1)
	}
	return ans
}
