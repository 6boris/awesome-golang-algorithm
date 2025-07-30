# [1805.Number of Different Integers in a String][title]

## Description
You are given a string `word` that consists of digits and lowercase English letters.

You will replace every non-digit character with a space. For example, `"a123bc34d8ef34"` will become `" 123  34 8  34"`. Notice that you are left with some integers that are separated by at least one space: `"123"`, `"34"`, `"8"`, and `"34"`.

Return the number of **different** integers after performing the replacement operations on `word`.

Two integers are considered different if their decimal representations **without any leading zeros** are different.

**Example 1:**

```
Input: word = "a123bc34d8ef34"
Output: 3
Explanation: The three different integers are "123", "34", and "8". Notice that "34" is only counted once.
```

**Example 2:**

```
Input: word = "leet1234code234"
Output: 2
```

**Example 3:**

```
Input: word = "a1b01c001"
Output: 1
Explanation: The three integers "1", "01", and "001" all represent the same integer because
the leading zeros are ignored when comparing their decimal values.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/number-of-different-integers-in-a-string/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
