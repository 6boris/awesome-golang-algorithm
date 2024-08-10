# [959.Regions Cut By Slashes][title]

## Description
An `n x n` grid is composed of `1 x 1` squares where each `1 x 1` square consists of a `'/'`, `'\'`, or blank space `' '`. These characters divide the square into contiguous regions.

Given the `grid` grid represented as a string array, return the number of regions.

Note that backslash characters are escaped, so a `'\'` is represented as `'\\'`.


**Example 1:**  

![1](./1.png)

```
Input: grid = [" /","/ "]
Output: 2
```

**Example 2:**

![2](./2.png)

```
Input: grid = [" /","  "]
Output: 1
```

**Example 3:**

![3](./4.png)

```
Input: grid = ["/\\","\\/"]
Output: 5
Explanation: Recall that because \ characters are escaped, "\\/" refers to \/, and "/\\" refers to /\.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/regions-cut-by-slashes/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
