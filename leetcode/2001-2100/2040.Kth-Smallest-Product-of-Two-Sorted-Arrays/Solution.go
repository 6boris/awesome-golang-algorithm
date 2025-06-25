package Solution

import (
	"math"
	"sort"
)

func Solution(nums1, nums2 []int, k int64) int64 {
	// 保证nums1更短，减少遍历次数
	if len(nums1) > len(nums2) {
		return Solution(nums2, nums1, k)
	}

	left, right := int64(-1e10), int64(1e10)
	for left < right {
		mid := left + (right-left)/2
		count := checkNums(mid, nums1, nums2)
		if count >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func checkNums(mid int64, nums1, nums2 []int) int64 {
	var ret int64 = 0
	for _, x := range nums1 {
		if x == 0 {
			if mid >= 0 {
				ret += int64(len(nums2))
			}
			// 如果mid<0，乘积都>=0，不计入
		} else if x > 0 {
			// 计算上界：mid / x
			yy := int64(math.Floor(float64(mid) / float64(x)))
			// upper_bound
			idx := sort.Search(len(nums2), func(i int) bool {
				return int64(nums2[i]) > yy
			})
			ret += int64(idx)
		} else {
			// x < 0
			// 计算下界：ceil(mid / x)
			yy := int64(math.Ceil(float64(mid) / float64(x)))
			// lower_bound
			idx := sort.Search(len(nums2), func(i int) bool {
				return int64(nums2[i]) >= yy
			})
			ret += int64(len(nums2) - idx)
		}
	}
	return ret
}
