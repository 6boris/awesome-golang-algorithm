package Solution

func Solution(startValue, target int) int {
	ans := 0
	for target > startValue {
		ans++
		if target&1 == 1 {
			target++
			continue
		}
		target /= 2
	}
	return ans + startValue - target
}
