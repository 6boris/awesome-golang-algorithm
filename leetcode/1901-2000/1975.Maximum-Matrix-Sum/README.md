# [1975.Maximum Matrix Sum][title]

## Description
You are given an `n x n` integer `matrix`. You can do the following operation **any** number of times:

- Choose any two **adjacent** elements of `matrix` and **multiply** each of them by ``-1`.

Two elements are considered **adjacent** if and only if they share a **border**.

Your goal is to **maximum** the summation of the matrix's elements. Return the **maximum** sum of the matrix's elements using the operation mentioned above.

**Example 1:**  

![1](./1.png)

```
Input: matrix = [[1,-1],[-1,1]]
Output: 4
Explanation: We can follow the following steps to reach sum equals 4:
- Multiply the 2 elements in the first row by -1.
- Multiply the 2 elements in the first column by -1.
```

**Example 2:**  

![2](./2.png)

```
Input: matrix = [[1,2,3],[-1,-2,-3],[1,2,3]]
Output: 16
Explanation: We can follow the following step to reach sum equals 16:
- Multiply the 2 last elements in the second row by -1.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/maximum-matrix-sum/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
