# [1339.Maximum Product of Splitted Binary Tree][title]

## Description
Given the `root` of a binary tree, split the binary tree into two subtrees by removing one edge such that the product of the sums of the subtrees is maximized.

Return the maximum product of the sums of the two subtrees. Since the answer may be too large, return it **modulo** 10<sup>9</sup> + 7.

**Note** that you need to maximize the answer before taking the mod and not after taking it.

**Example 1:**  
![example1](./sample_1_1699.png)

```
Input: root = [1,2,3,4,5,6]
Output: 110
Explanation: Remove the red edge and get 2 binary trees with sum 11 and 10. Their product is 110 (11*10)
```

**Example 2:**  
![example2](./sample_2_1699.png)

```
Input: root = [1,null,2,3,4,null,null,5,6]
Output: 90
Explanation: Remove the red edge and get 2 binary trees with sum 15 and 6.Their product is 90 (15*6)
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/maximum-product-of-splitted-binary-tree/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
