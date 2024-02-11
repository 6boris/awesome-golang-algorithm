# [2023.Number of Pairs of Strings With Concatenation Equal to Target][title]

## Description
Given an array of **digit** strings nums and a **digit** string `target`, return the number of pairs of indices `(i, j)` (where `i != j`) such that the **concatenation** of `nums[i] + nums[j]` equals `target`.

**Example 1:**

```
Input: nums = ["777","7","77","77"], target = "7777"
Output: 4
Explanation: Valid pairs are:
- (0, 1): "777" + "7"
- (1, 0): "7" + "777"
- (2, 3): "77" + "77"
- (3, 2): "77" + "77"
```

**Example 2:**

```
Input: nums = ["123","4","12","34"], target = "1234"
Output: 2
Explanation: Valid pairs are:
- (0, 1): "123" + "4"
- (2, 3): "12" + "34"
```

**Example 3:**

```
Input: nums = ["1","1","1"], target = "11"
Output: 6
Explanation: Valid pairs are:
- (0, 1): "1" + "1"
- (1, 0): "1" + "1"
- (0, 2): "1" + "1"
- (2, 0): "1" + "1"
- (1, 2): "1" + "1"
- (2, 1): "1" + "1"
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/number-of-pairs-of-strings-with-concatenation-equal-to-target/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
