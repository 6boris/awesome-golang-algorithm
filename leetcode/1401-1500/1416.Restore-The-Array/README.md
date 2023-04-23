# [1416. Restore The Array][title]

## Description
A program was supposed to print an array of integers. The program forgot to print whitespaces and the array is printed as a string of digits s and all we know is that all integers in the array were in the range `[1, k]` and there are no leading zeros in the array.

Given the string `s` and the integer `k`, return the number of the possible arrays that can be printed as `s` using the mentioned program. Since the answer may be very large, return it **modulo** `10^9 + 7`.

**Example 1:**

```
Input: s = "1000", k = 10000
Output: 1
Explanation: The only possible array is [1000]
```

**Example 2:**

```
Input: s = "1000", k = 10
Output: 0
Explanation: There cannot be an array that was printed this way and has all integer >= 1 and <= 10.
```

**Example 3:**

```
Input: s = "1317", k = 2000
Output: 8
Explanation: Possible arrays are [1317],[131,7],[13,17],[1,317],[13,1,7],[1,31,7],[1,3,17],[1,3,1,7]
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/restore-the-array/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
