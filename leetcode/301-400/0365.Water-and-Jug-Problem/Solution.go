package Solution

func gcd365(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
func Solution(x int, y int, target int) bool {
	if x+y < target {
		return false
	}
	g := gcd365(x, y)
	return target%g == 0
}
