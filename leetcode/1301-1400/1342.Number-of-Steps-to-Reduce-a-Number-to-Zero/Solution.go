package Solution

func Solution(num int) int {
	if num == 0 {
		return 0
	}
	if num == 1 {
		return 1
	}

	ans := Solution(num/2) + 1
	if num&1 == 1 {
		ans++
	}

	return ans
}
