# [838.Push Dominoes][title]

## Description
There are `n` dominoes in a line, and we place each domino vertically upright. In the beginning, we simultaneously push some of the dominoes either to the left or to the right.

After each second, each domino that is falling to the left pushes the adjacent domino on the left. Similarly, the dominoes falling to the right push their adjacent dominoes standing on the right.

When a vertical domino has dominoes falling on it from both sides, it stays still due to the balance of the forces.

For the purposes of this question, we will consider that a falling domino expends no additional force to a falling or already fallen domino.

You are given a string dominoes representing the initial state where:

- `dominoes[i] = 'L'`, if the i<sup>th</sup> domino has been pushed to the left,
- `dominoes[i] = 'R'`, if the i<sup>th</sup> domino has been pushed to the right, and
- `dominoes[i] = '.'`, if the i<sup>th</sup> domino has not been pushed.

Return a string representing the final state.

**Example 1:**

```
Input: dominoes = "RR.L"
Output: "RR.L"
Explanation: The first domino expends no additional force on the second domino.
```

**Example 2:**  

![1](./domino.png)

```
Input: dominoes = ".L.R...LR..L.."
Output: "LL.RR.LLRRLL.."
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/push-dominoes/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
