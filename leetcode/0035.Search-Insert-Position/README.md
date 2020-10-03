# [35. Search Insert Position][title]

## Description

Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.

**Example 1:**

```
Input: [1,3,5,6], 5
Output: 2
```

**Example 2:**

```
Input: [1,3,5,6], 2
Output: 1
```

**Example 3:**

```
Input: [1,3,5,6], 7
Output: 4
```

**Example 1:**

```
Input: [1,3,5,6], 0
Output: 0
```

**Tags:** Array, Binary Search


## 题解
### 思路1

- 直接二分查找小于等于目标值的第一个位置。
- 如果nums[l] >= target，说明找到了target或者数组中所有元素都比target小，则返回l
- 否则说明数组所有元素都比target大，此时l一定是0，故返回l。
时间复杂度

```go
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if target > nums[left] {
		left++
	}
	return left
}

```
### 思路2


## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/search-insert-position/description/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
