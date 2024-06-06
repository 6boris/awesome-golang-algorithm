package Solution

var (
	rotateMap = map[int]int{
		0: 0, 1: 1, 8: 8,
		2: 5, 5: 2, 6: 9, 9: 6,
	}
)

func isNumOk(n int) bool {
	ans := false
	for n > 0 {
		x := n % 10
		v, ok := rotateMap[x]
		if !ok {
			return false
		}
		ans = ans || x != v
		n /= 10
	}
	return ans
}
func Solution(n int) int {
	ans := 0
	for i := 1; i <= n; i++ {
		if isNumOk(i) {
			ans++
		}
	}
	return ans
}
