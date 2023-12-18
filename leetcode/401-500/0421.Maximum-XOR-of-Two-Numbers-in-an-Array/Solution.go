package Solution

/*
这个道题不可以用n^2的算法来做，数组的长度是2*10^5
所以就是构建最大xor的过程
*/
func Solution(nums []int) int {
	ans := 0
	mask := 0
	for i := 31; i >= 0; i-- {
		// i-31位的前缀表示的数字
		mask |= 1 << i
		set := make(map[int]struct{})
		for _, n := range nums {
			set[n&mask] = struct{}{}
		}
		// 尝试将ans第i位设置为1
		// 我们通过set得到了数组每一项的前缀。
		// 如果可以将第i位设置为1，那么就是说在set中存在两个数字 a, b使得 a^b=ans ans是我们的最大抑或值。
		// 抑或满足交换性，所以就是判断a^ans 是否在set中即可。如果在就可以将第i位设置为1。
		t := ans | (1 << i)
		for pre := range set {
			if _, ok := set[pre^t]; ok {
				ans = t
				break
			}
		}
	}

	return ans
}
