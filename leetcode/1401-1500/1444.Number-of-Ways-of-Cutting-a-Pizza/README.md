# [1444. Number of Ways of Cutting a Pizza][title]


## Description
Given a rectangular pizza represented as a `rows x cols` matrix containing the following characters: `'A'` (an apple) and `'.'` (empty cell) and given the integer `k`. You have to cut the pizza into `k` pieces using `k-1` cuts. 

For each cut you choose the direction: vertical or horizontal, then you choose a cut position at the cell boundary and cut the pizza into two pieces. If you cut the pizza vertically, give the left part of the pizza to a person. If you cut the pizza horizontally, give the upper part of the pizza to a person. Give the last piece of pizza to the last person.

Return the number of ways of cutting the pizza such that each piece contains **at least** one apple. Since the answer can be a huge number, return this modulo 10^9 + 7.

**Example 1:**  

![example1](./ways_to_cut_apple_1.png)

```
Input: pizza = ["A..","AAA","..."], k = 3
Output: 3 
Explanation: The figure above shows the three ways to cut the pizza. Note that pieces must contain at least one apple.
```

**Example 2:**

```
Input: pizza = ["A..","AA.","..."], k = 3
Output: 1
```

**Example 3:**

```
Input: pizza = ["A..","A..","..."], k = 1
Output: 1
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/number-of-ways-of-cutting-a-pizza/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
