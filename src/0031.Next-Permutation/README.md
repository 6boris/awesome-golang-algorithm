# [31. Next Permutation][title]

## Description

Given two binary strings, return their sum (also a binary string).

The input strings are both **non-empty** and contains only characters `1` or `0`.

**Example 1:**

```
Input: a = "11", b = "1"
Output: "100"
```

**Example 2:**

```
Input: a = "1010", b = "1011"
Output: "10101"
```

**Tags:** Math, String

## 题意
>判断下一个是否有满足要求的排列

## 题解

### 思路1
>  暴力全部扫描 时间复杂度可以飙到 O(N!),代码都不想写了

```go

```

### 思路2
> 官方解答
<div align="center">
    <img src="https://leetcode-cn.com/media/original_images/31/31_Next_Permutation.gif" alt="">
</div>

```go
func nextPermutation(nums []int) {
	i := len(nums) - 2

	for i >= 0 && nums[i+1] <= nums[i] {
		i--
	}
	if i >= 0 {
		j := len(nums) - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums, i+1)
}

func reverse(nums []int, start int) {
	i, j := start, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/next-permutation/description/
[me]: https://github.com/kylesliu/awesome-golang-leetcode
