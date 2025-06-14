# [2566.Maximum Difference by Remapping a Digit][title]

## Description

You are given an integer `num`. You know that Bob will sneakily **remap** one of the `10` possible digits (`0` to `9`) to another digit.

Return the difference between the maximum and minimum values Bob can make by remapping **exactly one** digit in `num`.

**Notes**:

- When Bob remaps a digit d1 to another digit d2, Bob replaces all occurrences of `d1` in `num` with `d2`.
- Bob can remap a digit to itself, in which case `num` does not change.
- Bob can remap different digits for obtaining minimum and maximum values respectively.
- The resulting number after remapping can contain leading zeroes.

**Example 1:**

```
Input: num = 11891
Output: 99009
Explanation: 
To achieve the maximum value, Bob can remap the digit 1 to the digit 9 to yield 99899.
To achieve the minimum value, Bob can remap the digit 1 to the digit 0, yielding 890.
The difference between these two numbers is 99009.
```

**Example 2:**

```
Input: num = 90
Output: 99
Explanation:
The maximum value that can be returned by the function is 99 (if 0 is replaced by 9) and the minimum value that can be returned by the function is 0 (if 9 is replaced by 0).
Thus, we return 99.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/maximum-difference-by-remapping-a-digit
[me]: https://github.com/kylesliu/awesome-golang-algorithm
