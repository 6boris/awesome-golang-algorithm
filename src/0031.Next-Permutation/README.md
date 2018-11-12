# [31. Next Permutation][title]

## Description

Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.

If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).

The replacement must be in-place and use only constant extra memory.

Here are some examples. Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.
**Example 1:**

```
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
```

**Example 2:**

```

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
