# [1. Add Sum][title]

## Problem

### Description

Given an array with n integers, your task is to check if it could become non-decreasing by modifying at most 1 element.

We define an array is non-decreasing if array[i] <= array[i + 1] holds for every i (1 <= i < n).
### Example
```cpp
Input: [4,2,3]
Output: True
Explanation: You could modify the first 4 to 1 to get a non-decreasing array.

Input: [4,2,1]
Output: False
Explanation: You can't get a non-decreasing array by modify at most one element.
```
> Note: 如果array [i] <= array [i + 1]对于每个i（1 <= i <n）成立，我们定义一个数组是非递减的。

>优先把 nums[i]降为nums[i+1]，这样可以减少 nums[i+1] > nums[i+2] 的风险。

## 题解
### 思路1
>题意是让你从给定的数组中找到两个元素的和为指定值的两个索引，最容易的当然是循环两次，复杂度为 O(n^2)，首次提交居然是 2ms，打败了 100% 的提交，谜一样的结果，之后再次提交就再也没跑到过 2ms 了。
```go
func twoSum(nums []int, target int) []int {
	if nums == nil || len(nums) < 2 {
		return []int{-1, -1}
	}

	for k, v := range nums {
		for i := k + 1; i < len(nums); i++ {
			if v+nums[i] == target {
				return []int{k, i}
			}
		}
	}
	return nil
}
```

### 思路2
>利用 HashMap 作为存储，键为目标值减去当前元素值，索引为值，比如 i = 0 时，此时首先要判断 nums[0] = 2 是否在 map 中，如果不存在，那么插入键值对 key = 9 - 2 = 7, value = 0，之后当 i = 1 时，此时判断 nums[1] = 7 已存在于 map 中，那么取出该 value = 0 作为第一个返回值，当前 i 作为第二个返回值，具体代码如下所示。
```go
func twoSum(nums []int, target int) []int {
	if nums == nil || len(nums) < 2 {
		return []int{-1, -1}
	}
	res := []int{-1, -1}

	//	MAP的KEY表示值，MAP的VAL表示nums的下标
	intMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		//	判断MAP中是否纯在一个Key满足 KEY + nums[i] = target
		//	如果满足则返回相关地址,不满足则将数组中的值PUSH到MAP
		if _, ok := intMap[target-nums[i]]; ok {
			res[0] = intMap[target-nums[i]]
			res[1] = i
			break
		}
		intMap[nums[i]] = i
	}
	return res
}
```


## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome golang leetcode][me]

[title]: https://leetcode.com/problems/two-sum/description/
[me]: https://github.com/kylesliu/awesome-golang-leetcode
