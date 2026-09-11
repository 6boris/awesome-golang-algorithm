# [3483.Unique 3-Digit Even Numbers][title]

## Description
You are given an array of digits called `digits`. Your task is to determine the number of **distinct** three-digit even numbers that can be formed using these digits.

**Note**: Each copy of a digit can only be used **once per number**, and there may **not** be leading zeros.

**Example 1:**

```
Input: digits = [1,2,3,4]

Output: 12

Explanation: The 12 distinct 3-digit even numbers that can be formed are 124, 132, 134, 142, 214, 234, 312, 314, 324, 342, 412, and 432. Note that 222 cannot be formed because there is only 1 copy of the digit 2.
```

**Example 2:**

```
Input: digits = [0,2,2]

Output: 2

Explanation: The only 3-digit even numbers that can be formed are 202 and 220. Note that the digit 2 can be used twice because it appears twice in the array.
```

**Example 3:**

```
Input: digits = [6,6,6]

Output: 1

Explanation: Only 666 can be formed.
```

**Example 4:**

```
Input: digits = [1,3,5]

Output: 0

Explanation: No even 3-digit numbers can be formed.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/unique-3-digit-even-numbers/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
