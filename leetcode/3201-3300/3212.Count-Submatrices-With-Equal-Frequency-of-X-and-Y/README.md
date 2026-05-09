# [3212.Count Submatrices With Equal Frequency of X and Y][title]

## Description
Given a 2D character matrix `grid`, where `grid[i][j]` is either `'X'`, `'Y'`, or `'.'`, return the number of submatrices that contain:

- `grid[0][0]`
- an **equal** frequency of `'X'` and `'Y'`.
- **at least** one `'X'`.

**Example 1:**  

![1](./1.png)

```
Input: grid = [["X","Y","."],["Y",".","."]]

Output: 3

Explanation:
```

**Example 2:**

```
Input: grid = [["X","X"],["X","Y"]]

Output: 0

Explanation:

No submatrix has an equal frequency of 'X' and 'Y'.
```

**Example 3:**

```
Input: grid = [[".","."],[".","."]]

Output: 0

Explanation:

No submatrix has at least one 'X'.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/count-submatrices-with-equal-frequency-of-x-and-y/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
