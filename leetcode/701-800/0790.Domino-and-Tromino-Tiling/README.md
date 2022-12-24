# [790.Domino and Tromino Tiling][title]

## Description
You have two types of tiles: a 2 x 1 domino shape and a tromino shape. You may rotate these shapes.  

![shape](./lc-domino.jpg)

Given an integer n, return the number of ways to tile an `2 x n` board. Since the answer may be very large, return it **modulo** 10<sup>9</sup> + 7.

In a tiling, every square must be covered by a tile. Two tilings are different if and only if there are two 4-directionally adjacent cells on the board such that exactly one of the tilings has both squares occupied by a tile.

**Example 1:**  

![example1](./lc-domino1.jpg)

```
Input: n = 3
Output: 5
Explanation: The five different ways are show above.
```


**Example 2:**

```
Input: n = 1
Output: 1
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/domino-and-tromino-tiling/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
