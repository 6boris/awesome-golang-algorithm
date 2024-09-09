# [1598.Crawler Log Folder][title]

## Description
The Leetcode file system keeps a log each time some user performs a change folder operation.

The operations are described below:

- `"../"` : Move to the parent folder of the current folder. (If you are already in the main folder, **remain in the same folder**).
- `"./"` : Remain in the same folder.
- `"x/"` : Move to the child folder named x (This folder is **guaranteed to always exist**).

You are given a list of strings `logs` where `logs[i]` is the operation performed by the user at the i<sup>th</sup> step.

The file system starts in the main folder, then the operations in `logs` are performed.

Return the minimum number of operations needed to go back to the main folder after the change folder operations.

**Example 1:**  

![1](./sample_11_1957.png)

```
Input: logs = ["d1/","d2/","../","d21/","./"]
Output: 2
Explanation: Use this change folder operation "../" 2 times and go back to the main folder.
```

**Example 2:**  

![2](./sample_22_1957.png)

```
Input: logs = ["d1/","d2/","./","d3/","../","d31/"]
Output: 3
```

**Example 3:**

```
Input: logs = ["d1/","../","../","../"]
Output: 0
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/crawler-log-folder/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
