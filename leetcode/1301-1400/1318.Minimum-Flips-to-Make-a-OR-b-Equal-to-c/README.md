# [1318.Minimum Flips to Make a OR b Equal to c][title]

## Description
Given 3 positives numbers `a`, `b` and `c`. Return the minimum flips required in some bits of `a` and `b` to make ( `a` OR `b` == `c` ). (bitwise OR operation).
Flip operation consists of change **any** single bit 1 to 0 or change the bit 0 to 1 in their binary representation.

**Example 1:**  

![example1](./sample_3_1676.png)

```
Input: a = 2, b = 6, c = 5
Output: 3
Explanation: After flips a = 1 , b = 4 , c = 5 such that (a OR b == c)
```

**Example 2:**

```
Input: a = 4, b = 2, c = 7
Output: 1
```

**Example 3:**

```
Input: a = 1, b = 2, c = 3
Output: 0
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/minimum-flips-to-make-a-or-b-equal-to-c/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
