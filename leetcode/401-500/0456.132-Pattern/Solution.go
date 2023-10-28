package Solution

import "sort"

// 借鉴了题解，本想通过排序部分+二分会超时.
func Solution(nums []int) bool {
	ll := len(nums)
	if ll < 3 {
		return false
	}
	minArr := make([]int, ll)
	minArr[0] = nums[0]
	for i := 1; i < ll; i++ {
		minArr[i] = nums[i]
		if minArr[i-1] < minArr[i] {
			minArr[i] = minArr[i-1]
		}
	}
	//var bsearch func(int, int, int) int
	/*
		bsearch = func(l, r, target int) int {
			if l == r || target < nums[l] {
				return -1
			}
			for l < r {
				mid := r - (r-l)/2
				if mid >= ll {
					break
				}
				if nums[mid] < target {
					l = mid
					continue
				}
				r = mid - 1
			}
			return l
		}
	*/
	// 搜索第一个大于等于target的下表
	/*
		bsearch = func(l, r, target int) int {

			for l <= r {
				mid := (l + r) / 2
				if nums[mid] >= target {
					r = mid - 1
					continue
				}
				l = mid + 1
			}
			return l
		}
	*/
	k := ll
	for j := ll - 1; j >= 0; j-- {
		if minArr[j] >= nums[j] {
			continue
		}
		/*
			sort.Ints(nums[j+1:])

			index := bsearch(j+1, ll, nums[j])
			if index != -1 && nums[index] > minArr[j] {
				return true
			}
		*/
		index := sort.Search(ll, func(i int) bool {
			if i >= j+1 && i < ll && nums[i] >= minArr[j]+1 {
				return true
			}
			return false
		})
		k = index - 0
		if k < ll && nums[k] < nums[j] {
			return true
		}
		k--
		nums[k] = nums[j]
	}
	return false
}
