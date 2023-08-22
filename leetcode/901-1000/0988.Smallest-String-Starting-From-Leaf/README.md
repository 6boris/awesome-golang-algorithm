# [988.Smallest String Starting From Leaf][title]

## Description
You are given the `root` of a binary tree where each node has a value in the range `[0, 25]` representing the letters `'a'` to `'z'`.

Return the **lexicographically smallest** string that starts at a leaf of this tree and ends at the root.

As a reminder, any shorter prefix of a string is **lexicographically smaller**.

- For example, `"ab"` is lexicographically smaller than `"aba"`.

A leaf of a node is a node that has no children.

**Example 1:**  

![example1](./tree1.png)

```
Input: root = [0,1,2,3,4,3,4]
Output: "dba"
```

**Example 2:**  

![example2](./tree2.png)

```
Input: root = [25,1,3,1,3,0,2]
Output: "adz"
```

**Example 3:**  

![example3](./tree3.png)

```
Input: root = [2,2,1,null,1,0,null,0]
Output: "abc"
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/smallest-string-starting-from-leaf/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
