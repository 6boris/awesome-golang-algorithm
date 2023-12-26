# [447.Number of Boomerangs][title]

## Description
You are given `n` `points` in the plane that are all **distinct**, where `points[i] = [xi, yi]`. A **boomerang** is a tuple of points `(i, j, k)` such that the distance between `i` and `j` equals the distance between `i` and `k` (**the order of the tuple matters**).

Return the number of boomerangs.

**Example 1:**

```
Input: points = [[0,0],[1,0],[2,0]]
Output: 2
Explanation: The two boomerangs are [[1,0],[0,0],[2,0]] and [[1,0],[2,0],[0,0]].
```

**Example 2:**

```
Input: points = [[1,1],[2,2],[3,3]]
Output: 2
```

**Example 3:**

```
Input: points = [[1,1]]
Output: 0
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/number-of-boomerangs/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
