package Solution

// 想要将n清理，就需要把每个位上的数字都变成0，而且这个只能由操作deci-binary数字
// 所以就是找到哪个位上的数字最大， 就是清零的次数
func Solution(n string) int {
	ans := 0
	for _, b := range []byte(n) {
		if int(b-'0') > ans {
			ans = int(b - '0')
		}
	}
	return ans
}
