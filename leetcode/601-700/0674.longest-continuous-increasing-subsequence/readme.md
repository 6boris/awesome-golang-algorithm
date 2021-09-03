# [674.  Longest Continuous Increasing Subsequence][title]

## Description

Given an unsorted array of integers, find the length of longest continuous increasing subsequence (subarray).

**Example 1:**
```
Input: [1,3,5,4,7]
Output: 3
Explanation: The longest continuous increasing subsequence is [1,3,5], its length is 3. 
Even though [1,3,5,7] is also an increasing subsequence, it's not a continuous one where 5 and 7 are separated by 4. 
```

**Example 2:**
```
Input: [2,2,2,2,2]
Output: 1
Explanation: The longest continuous increasing subsequence is [2], its length is 1. 
```

## Note
Length of the array will not exceed 10,000.

## 题意
> 求一个数组的最长连续上升序列长度。

## 题解

### 思路 1
> 判断相邻两个元素，后一个是否大于前一个，如果是就 +1，否则重置计数器。

```go
func findLengthOfLCIS(nums []int) int {
    if 0 == len(nums) {
        return 0
    }
    ans, sum := 1, 1
    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[i-1] {
            sum++
        } else {
            sum = 1
        }
        if sum > ans {
            ans = sum
        }
    }
    return ans
}
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/longest-continuous-increasing-subsequence/description/  
[me]: https://github.com/kylesliu/awesome-golang-algorithm
