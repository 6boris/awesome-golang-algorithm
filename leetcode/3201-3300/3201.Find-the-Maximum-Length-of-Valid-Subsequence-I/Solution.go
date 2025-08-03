package Solution

func Solution(nums []int) int {
	// 要么一连串偶数(奇数)，可以保证%2都是0
	a, b := 0, 0
	for _, n := range nums {
		if n&1 == 0 {
			a++
			continue
		}
		b++
	}
	ans := max(a, b)
	// 开始找一个奇数一个偶数, 保证了结果一定是奇数
	evenLen, oddLen := 0, 0
	for _, n := range nums {
		if n&1 == 0 {
			evenLen = oddLen + 1
			continue
		}
		oddLen = evenLen + 1
	}
	return max(ans, oddLen, evenLen)
}
