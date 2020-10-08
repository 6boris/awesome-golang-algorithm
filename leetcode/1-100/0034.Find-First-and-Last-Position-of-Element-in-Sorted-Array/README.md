# [34. Find First and Last Position of Element in Sorted Array][title]

## Description

Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].

**Example 1:**

```
Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
```

**Example 2:**

```
Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
```

**Tags:** Math, String

## 题意
>给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置
你的算法时间复杂度必须是 O(log n) 级别。
如果数组中不存在目标值，返回 [-1, -1]。
## 题解

### 思路1
>直线查找

```go
unc searchRange(nums []int, target int) []int {
	targetRange := []int{-1, -1}

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			targetRange[0] = i
			break
		}
	}
	if targetRange[0] == -1 {
		return targetRange
	}

	for j := len(nums) - 1; j >= 0; j-- {
		if nums[j] == target {
			targetRange[1] = j
			break
		}
	}
	return targetRange
}
```

### 思路2
> 二分查找
```go
func searchRange2(nums []int, target int) []int {
	targetRange := []int{-1, -1}
	leftIndex := extremeInsertionIndex(nums, target, true)

	if leftIndex == len(nums) || nums[leftIndex] != target {
		return targetRange
	}

	targetRange[0] = leftIndex
	targetRange[1] = extremeInsertionIndex(nums, target, false) - 1

	return targetRange

}
func extremeInsertionIndex(nums []int, target int, left bool) int {
	lo := 0
	hi := len(nums)
	for lo < hi {
		mid := (lo + hi) / 2
		if nums[mid] > target || left && target == nums[mid] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
