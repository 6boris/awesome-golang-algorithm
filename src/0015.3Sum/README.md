# [15. 3Sum][title]


## Description

Given an array `nums` of *n* integers, are there elements *a*, *b*, *c* in `nums` such that *a* + *b* + *c* = 0? Find all unique triplets in the array which gives the sum of zero.

**Note:**

The solution set must not contain duplicate triplets.

**Example:**

```
Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
```

**Tags:** Array, Two Pointers

## 题意
>从数组中找出所有三个和为 0 的元素构成的非重复序列

## 题解

### 思路1
> 我们可以把数组先做下排序，然后遍历这个排序数组，用两个指针分别指向当前元素的下一个和数组尾部，判断三者的和与 0 的大小来移动两个指针，其中细节操作就是优化和去重。

```go
func threeSum(nums []int) [][]int {
	ans := [][]int{}
	if len(nums) <= 2 {
		return ans
	}

	//	排序
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			low, high, sum := i+1, len(nums)-1, 0-nums[i]
			for low < high {
				//	和为0的时候
				if nums[low]+nums[high] == sum {
					ans = append(ans, []int{nums[i], nums[low], nums[high]})
					//	排除；连续相等的数
					for low < high && nums[low] == nums[low+1] {
						low++
					}
					for low < high && nums[high] == nums[high-1] {
						high--
					}
					low++
					high--
				} else if nums[low]+nums[high] < sum {
					low++
				} else {
					high--
				}

			}

		}
	}
	return ans
}
```

### 思路2
> 思路2
```go

```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/3sum/description/
[me]: https://github.com/kylesliu/awesome-golang-leetcode
