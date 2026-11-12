# [3945.Digit Frequency Score][title]

## Description
You are given an integer `n`.

The **score** of n is defined as the **sum** of `d * freq(d)` over all **distinct** digits `d`, where `freq(d)` denotes the number of times the digit `d` appears in `n`.

Return an integer denoting the score of `n`.

**Example 1:**

```
Input: n = 122

Output: 5

Explanation:

The digit 1 appears 1 time, contributing 1 * 1 = 1.
The digit 2 appears 2 times, contributing 2 * 2 = 4.
Thus, the score of n is 1 + 4 = 5.
```

**Example 2:**

```
Input: n = 101

Output: 2

Explanation:

The digit 0 appears 1 time, contributing 0 * 1 = 0.
The digit 1 appears 2 times, contributing 1 * 2 = 2.
Thus, the score of n is 2.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/digit-frequency-score/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
