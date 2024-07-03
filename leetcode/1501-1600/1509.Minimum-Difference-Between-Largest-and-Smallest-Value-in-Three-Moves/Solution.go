package Solution

import (
	"math"
	"sort"
)

func Solution(nums []int) int {
	l := len(nums)
	if l <= 4 {
		return 0
	}
	sort.Ints(nums)
	// 如果最后元素全都变了，那就是nums[l-4]-num[0]了
	// 如果不是处理最后三个元素而至最后两个，然后把最小的给干掉，将最小的数值提升，然后看
	// 0,1,1,    4, 6,      6, 6, 6
	// 如果处理最后三个最大的就是nums[l-4], 然后是nums[l-3], ...
	// 如果处理后面两个，那就是提升最小的值，
	// ... 如果最后面的数据一个都不处理，那就是直接提升最小值
	left := 0
	ans := math.MaxInt
	for i := l - 4; i < l; i++ {
		ans = min(ans, nums[i]-nums[left])
		left++
	}

	return ans
	//return nums[len(nums)-4] - nums[0]
}
