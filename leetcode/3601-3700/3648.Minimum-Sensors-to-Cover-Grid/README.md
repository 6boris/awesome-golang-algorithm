# [3648.Minimum Sensors to Cover Grid][title]

## Description
You are given `n × m` grid and an integer `k`.

A sensor placed on cell (`r, c`) covers all cells whose **Chebyshev distance** from (`r, c`) is at most `k`.

The **Chebyshev distance** between two cells (`r1, c1`) and (`r2, c2`) is `max(|r1 − r2|,|c1 − c2|)`.

Your task is to return the **minimum** number of sensors required to cover every cell of the grid.

**Example 1:**

```
Input: n = 5, m = 5, k = 1

Output: 4

Explanation:

Placing sensors at positions (0, 3), (1, 0), (3, 3), and (4, 1)
```

**Example 2:**

```
Input: n = 2, m = 2, k = 2

Output: 1

Explanation:

With k = 2, a single sensor can cover the entire 2 * 2 grid regardless of its position. Thus, the answer is 1.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/minimum-sensors-to-cover-grid/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
