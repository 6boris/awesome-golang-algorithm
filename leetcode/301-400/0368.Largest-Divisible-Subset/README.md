# [368.Largest Divisible Subset][title]

## Description
Given a set of __distinct__ positive integers nums, return the largest subset `answer` such that every pair `(answer[i], answer[j])` of elements in this subset satisfies:
- `answer[i] % answer[j] == 0`, or
- `answer[j] % answer[i] == 0`

If there are multiple solutions, return any of them.

**Example 1:**

```
Input: nums = [1,2,3]
Output: [1,2]
Explanation: [1,3] is also accepted.
```

__Example 2:__

```
Input: nums = [1,2,4,8]
Output: [1,2,4,8]
```


## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/largest-divisible-subset/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
