# [3963.Create Grid With Exactly One Path][title]

## Description
You are given two integers `m` and `n`, representing the number of rows and columns of a grid.

Construct **any** `m x n` grid consisting only of the characters `'.'` and `'#'`, where:

- `'.'` represents a free cell.
- `'#'` represents an obstacle cell.

A **valid path** is a sequence of free cells that:

- Starts at the top-left cell `(0, 0)`.
- Ends at the bottom-right cell `(m - 1, n - 1)`.
- Moves only:

    - Right, from `(i, j)` to `(i, j + 1)`, or
    - Down, from `(i, j)` to `(i + 1, j)`.

Return any grid such that there is **exactly one valid path** from the top-left cell to the bottom-right cell.

**Example 1:**  

![1](./1.png)

```
Input: m = 2, n = 3

Output: ["..#","#.."]

Explanation:

The only valid path is: (0,0) → (0,1) → (1,1) → (1,2)
```

**Example 2:**  

![2](./2.png)

```
Input: m = 3, n = 3

Output: ["..#","#..","##."]

Explanation:

The only valid path is: (0,0) → (0,1) → (1,1) → (1,2) → (2,2)
```

**Example 3:**

```
Input: m = 1, n = 4

Output: ["...."]

Explanation:

The only valid path is: (0,0) → (0,1) → (0,2) → (0,3)
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/create-grid-with-exactly-one-path/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
