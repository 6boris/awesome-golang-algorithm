# [3918.Sum of Primes Between Number and Its Reverse][title]

## Description
You are given an integer `n`.

Let `r` be the integer formed by reversing the digits of `n`.

Return the **sum** of all prime numbers between `min(n, r)` and `max(n, r)`, inclusive.

**Example 1:**

```
Input: n = 13

Output: 132

Explanation:

The reverse of 13 is 31. Thus, the range is [13, 31].
The prime numbers in this range are 13, 17, 19, 23, 29, and 31.
The sum of these prime numbers is 13 + 17 + 19 + 23 + 29 + 31 = 132.
```

**Example 2:**

```
Input: n = 10

Output: 17

Explanation:

The reverse of 10 is 1. Thus, the range is [1, 10].
The prime numbers in this range are 2, 3, 5, and 7.
The sum of these prime numbers is 2 + 3 + 5 + 7 = 17.
```

**Example 3:**

```
Input: n = 8

Output: 0

Explanation:

The reverse of 8 is 8. Thus, the range is [8, 8].
There are no prime numbers in this range, so the sum is 0.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/sum-of-primes-between-number-and-its-reverse/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
