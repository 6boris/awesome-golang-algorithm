# [18. 4Sum][title]

## Description

Given an array `nums` of *n* integers and an integer `target`, are there elements *a*, *b*, *c*, and *d* in `nums` such that *a* + *b* + *c* + *d* = `target`? Find all unique quadruplets in the array which gives the sum of `target`.

**Note:**

The solution set must not contain duplicate quadruplets.

**Example:**

```
Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.

A solution set is:
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
```

**Tags:** Array, Hash Table, Two Pointers

## 题意
>题意是让你从数组中找出所有四个数的和为 `target` 的元素构成的非重复序列

## 题解

### 思路1
> 该题和  的思路基本一样，先对数组进行排序，然后遍历这个排序数组，因为这次是四个元素的和，所以外层需要两重循环，然后还是用两个指针分别指向当前元素的下一个和数组尾部，判断四者的和与 `target` 的大小来移动两个指针，其中细节操作还是优化和去重。

```go
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-3; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		threeSum(&res, nums[i], nums[i+1:], target-nums[i])
	}
	return res
}

func threeSum(res *[][]int, first int, nums []int, target int) {
	nlen := len(nums)
	for i := 0; i < nlen-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, nlen-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				*res = append(*res, []int{first, nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < target {
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++
			} else {
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				right--
			}
		}
	}
}
```

### 思路2
> 思路2
```go

```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/4sum/description/
[me]: https://github.com/kylesliu/awesome-golang-leetcode
