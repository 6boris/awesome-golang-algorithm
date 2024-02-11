package Solution

import "sort"

const mod1818 = 1000000007

func Solution(nums1 []int, nums2 []int) int {
	next := make([]int, len(nums1))
	copy(next, nums1)
	sort.Ints(next)
	shouldCut := -1
	var bsearch func(cur, target int) int
	bsearch = func(cur, target int) int {
		ans := -1
		left, right := 0, len(next)
		for left < right {
			mid := (right + left) / 2
			diff := next[mid] - cur
			if diff > 0 {
				if diff < target && (ans == -1 || ans > diff) {
					ans = diff
				}
				right = mid
			} else {
				diff = -diff
				if diff < target && (ans == -1 || ans > diff) {
					ans = diff
				}
				left = mid + 1
			}
		}
		return ans
	}
	diff := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		diff[i] = nums1[i] - nums2[i]
		if diff[i] < 0 {
			diff[i] = -diff[i]
		}
		if diff[i] == 0 {
			continue
		}
		r := bsearch(nums2[i], diff[i])
		if r != -1 && (shouldCut == -1 || diff[i]-r > shouldCut) {
			shouldCut = diff[i] - r
		}
	}

	ans := 0
	if shouldCut != -1 {
		ans = -shouldCut
	}

	for i := 0; i < len(diff); i++ {
		if ans < 0 {
			ans += diff[i]
			continue
		}
		ans = (ans + diff[i]) % mod1818
	}
	return ans
}
