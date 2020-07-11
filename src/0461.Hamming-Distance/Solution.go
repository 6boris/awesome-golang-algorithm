package Solution

func hammingDistance(x int, y int) int {
	ans, z := 0, x^y
	for z > 0 {
		ans += z & 1
		z >>= 1
	}
	return ans
}
