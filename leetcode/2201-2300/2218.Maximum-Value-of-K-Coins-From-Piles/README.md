# [2218.Maximum Value of K Coins From Piles][title]

## Description
There are n **piles** of coins on a table. Each pile consists of a **positive number** of coins of assorted denominations.

In one move, you can choose any coin on **top** of any pile, remove it, and add it to your wallet.

Given a list `piles`, where `piles[i]` is a list of integers denoting the composition of the i<sup>th</sup> pile from **top to bottom**, and a positive integer `k`, 
return the **maximum total value** of coins you can have in your wallet if you choose **exactly** `k` coins optimally.

**Example 1:**  

![example1](./e1.png)

```
Input: piles = [[1,100,3],[7,8,9]], k = 2
Output: 101
Explanation:
The above diagram shows the different ways we can choose k coins.
The maximum total we can obtain is 101.
```

**Example 2:**

```
Input: piles = [[100],[100],[100],[100],[100],[100],[1,1,1,1,1,1,700]], k = 7
Output: 706
Explanation:
The maximum total can be obtained if we choose all coins from the last pile.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/maximum-value-of-k-coins-from-piles/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
