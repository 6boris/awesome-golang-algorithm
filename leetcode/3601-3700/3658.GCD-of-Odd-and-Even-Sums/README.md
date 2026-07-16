# [3658.GCD of Odd and Even Sums][title]

## Description
You are given an integer `n`. Your task is to compute the **GCD** (greatest common divisor) of two values:

- `sumOdd`: the sum of the smallest `n` positive odd numbers.
- `sumEven`: the sum of the smallest `n` positive even numbers.

Return the GCD of `sumOdd` and `sumEven`.

**Example 1:**

```
Input: n = 4

Output: 4

Explanation:

Sum of the first 4 odd numbers sumOdd = 1 + 3 + 5 + 7 = 16
Sum of the first 4 even numbers sumEven = 2 + 4 + 6 + 8 = 20
Hence, GCD(sumOdd, sumEven) = GCD(16, 20) = 4.
```

**Example 2:**

```
Input: n = 5

Output: 5

Explanation:

Sum of the first 5 odd numbers sumOdd = 1 + 3 + 5 + 7 + 9 = 25
Sum of the first 5 even numbers sumEven = 2 + 4 + 6 + 8 + 10 = 30
Hence, GCD(sumOdd, sumEven) = GCD(25, 30) = 5.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/gcd-of-odd-and-even-sums/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
