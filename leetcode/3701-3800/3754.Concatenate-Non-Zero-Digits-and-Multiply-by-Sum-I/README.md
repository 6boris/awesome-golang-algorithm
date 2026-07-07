# [3754.Concatenate Non-Zero Digits and Multiply by Sum I][title]

## Description
You are given an integer `n`.

Form a new integer `x` by concatenating all the **non-zero digits** of `n` in their original order. If there are no **non-zero digits**, `x = 0`.

Let `sum` be the sum of digits in `x`.

Return an integer representing the value of `x * sum`.


**Example 1:**

```
Input: n = 10203004

Output: 12340

Explanation:

The non-zero digits are 1, 2, 3, and 4. Thus, x = 1234.
The sum of digits is sum = 1 + 2 + 3 + 4 = 10.
Therefore, the answer is x * sum = 1234 * 10 = 12340.
```

**Example 2:**

```
Input: n = 1000

Output: 1

Explanation:

The non-zero digit is 1, so x = 1 and sum = 1.
Therefore, the answer is x * sum = 1 * 1 = 1.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/concatenate-non-zero-digits-and-multiply-by-sum-i/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
