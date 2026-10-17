package Solution

import "sort"

func Solution(nums []int) int {
	n := len(nums)
	src := make([][2]int, n)
	var sum int
	for i := range nums {
		sum = 0
		for x := nums[i]; x > 0; x /= 10 {
			sum += x % 10
		}
		src[i] = [2]int{nums[i], sum}
	}
	dst := make([][2]int, n)
	copy(dst, src)
	sort.Slice(dst, func(i, j int) bool {
		a, b := dst[i], dst[j]
		if a[1] == b[1] {
			return a[0] < b[0]
		}
		return a[1] < b[1]
	})

	// 一个环的最小交换次数就是n-1
	// 乱序中存在一个多个小环的情况, 比如，目标是1，2，3，4，5，6
	// 但是实际1，2，3乱了，4，5，6乱了，1，2，3自己交换即可，同理4，5，6自己交:
	// 开始寻找环
	targetIndex := make(map[int]int)
	for i := range dst {
		targetIndex[dst[i][0]] = i
	}
	seen := make([]bool, n)
	var ret, cnt int
	for i := range nums {
		if nums[i] == dst[i][0] || seen[i] {
			continue
		}
		// nums[i]寻找合适的地方
		cnt = 0
		for !seen[i] {
			seen[i] = true
			cnt++
			i = targetIndex[nums[i]]
		}
		ret += cnt - 1
	}
	return ret
}
