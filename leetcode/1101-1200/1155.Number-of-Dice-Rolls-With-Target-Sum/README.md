# [1155.Number of Dice Rolls With Target Sum][title]

## Description
You have `n` dice, and each die has `k` faces numbered from `1` to `k`.

Given three integers `n`, `k`, and `target`, return the number of possible ways (out of the k<sup>n</sup> total ways) to roll the dice, so the sum of the face-up numbers equals `target`. Since the answer may be too large, return it modulo `10^9 + 7`.

**Example 1:**

```
Input: n = 1, k = 6, target = 3
Output: 1
Explanation: You throw one die with 6 faces.
There is only one way to get a sum of 3.
```

**Example 2:**

```
Input: n = 2, k = 6, target = 7
Output: 6
Explanation: You throw two dice, each with 6 faces.
There are 6 ways to get a sum of 7: 1+6, 2+5, 3+4, 4+3, 5+2, 6+1.
```

**Example 3:**

```
Input: n = 30, k = 30, target = 500
Output: 222616187
Explanation: The answer must be returned modulo 109 + 7.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/number-of-dice-rolls-with-target-sum/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
