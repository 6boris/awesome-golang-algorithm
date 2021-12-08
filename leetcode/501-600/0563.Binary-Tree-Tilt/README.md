# [563.Binary Tree Tilt][title]

## Description
Given the `root` of a binary tree, return the sum of every tree node's __tilt__.

The __tilt__ of a tree node is the __absolute difference__ between the sum of all left subtree node __values__ and all right subtree node __values__. If a node does not have a left child, then the sum of the left subtree node __values__ is treated as `0`. The rule is similar if there the node does not have a right child.


**Example 1:**  
![tilt1](./tilt1.jpeg)

```
Input: root = [1,2,3]
Output: 1
Explanation: 
Tilt of node 2 : |0-0| = 0 (no children)
Tilt of node 3 : |0-0| = 0 (no children)
Tilt of node 1 : |2-3| = 1 (left subtree is just left child, so sum is 2; right subtree is just right child, so sum is 3)
Sum of every tilt : 0 + 0 + 1 = 1
```

**Example 2:**  
![tilt2](./tilt2.jpeg)
```
Input: root = [4,2,9,3,5,null,7]
Output: 15
Explanation: 
Tilt of node 3 : |0-0| = 0 (no children)
Tilt of node 5 : |0-0| = 0 (no children)
Tilt of node 7 : |0-0| = 0 (no children)
Tilt of node 2 : |3-5| = 2 (left subtree is just left child, so sum is 3; right subtree is just right child, so sum is 5)
Tilt of node 9 : |0-7| = 7 (no left child, so sum is 0; right subtree is just right child, so sum is 7)
Tilt of node 4 : |(3+5+2)-(9+7)| = |10-16| = 6 (left subtree values are 3, 5, and 2, which sums to 10; right subtree values are 9 and 7, which sums to 16)
Sum of every tilt : 0 + 0 + 0 + 2 + 7 + 6 = 15
```

**Example 3:**  
![tilt3](./tilt3.jpeg)
```
Input: root = [21,7,14,1,1,2,2,3,3]
Output: 9
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/binary-tree-tilt/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
