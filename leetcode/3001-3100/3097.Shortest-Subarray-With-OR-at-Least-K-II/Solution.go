package Solution

func Solution(nums []int, k int) int {
	// 统计一个范围的1的位置的个数, 直接ok操作，没有个数统计，应该是没法滑动窗口，我没想到怎么滑动
	// 转成数字看是否比k大
	bitCount := [32]int{}
	var (
		setBit   func(int, int)
		toNumber func() int
	)
	setBit = func(n, del int) {
		one := 1
		for i := 0; i < 32; i++ {
			if n&one == one {
				bitCount[i] += del
			}
			one <<= 1
		}
	}
	toNumber = func() int {
		res := 0
		cur := 1
		for i := 0; i < 32; i++ {
			if bitCount[i] > 0 {
				res += cur
			}
			cur <<= 1
		}
		return res
	}
	start, end := 0, 0
	ans := -1
	for ; end < len(nums); end++ {
		setBit(nums[end], 1)
		r := toNumber()
		if r < k {
			continue
		}
		for ; start < end; start++ {
			setBit(nums[start], -1)
			if toNumber() < k {
				setBit(nums[start], 1)
				break
			}
		}
		tmpL := end - start + 1
		if ans == -1 || tmpL < ans {
			ans = tmpL
		}
	}
	return ans
}
