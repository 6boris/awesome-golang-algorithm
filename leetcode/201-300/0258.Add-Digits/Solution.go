package Solution

// 借鉴官方解题思路
func Solution(n int) int {
	if n == 0 {
		return 0
	}
	r := n % 9
	if r == 0 {
		return 9
	}
	return r
}
