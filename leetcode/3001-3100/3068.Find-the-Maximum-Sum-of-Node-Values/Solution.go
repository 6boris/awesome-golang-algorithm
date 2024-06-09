package Solution

import "sort"

func Solution(nums []int, k int, edges [][]int) int64 {
	// 经过测试，[0,1],[1,2] [1,2,1]这组数据也是6，就是说他选的任何节点俩节点，都可以
	// 那么我们关注的就是，最大值从小大排列， 而且题目保证了是一棵树，
	// 所以，只需要每次选俩和是正的就ok

	diff := make([]int, len(nums))
	ans := int64(0)
	for i, v := range nums {
		ans += int64(v)
		diff[i] = v ^ k - v
	}

	sort.Slice(diff, func(i, j int) bool {
		return diff[i] > diff[j]
	})
	for i := 0; i < len(diff); i += 2 {
		if i+1 == len(diff) {
			break
		}
		if s := int64(diff[i]) + int64(diff[i+1]); s > 0 {
			ans += s
		}
	}
	return ans
}
