# [997.Find the Town Judge][title]

## Description
In a town, there are `n` people labeled from `1` to `n`. There is a rumor that one of these people is secretly the town judge.

If the town judge exists, then:  

1. The town judge trusts nobody.
2. Everybody (except for the town judge) trusts the town judge.
3. There is exactly one person that satisfies properties **1** and **2**.  

You are given an array `trust` where `trust[i] = [ai, bi]` representing that the person labeled a<sub>i</sub> trusts the person labeled b<sub>i</sub>.

Return the label of the town judge if the town judge exists and can be identified, or return `-1` otherwise.  


**Example 1:**

```
Input: n = 2, trust = [[1,2]]
Output: 2
```

**Example 2:**
```
Input: n = 3, trust = [[1,3],[2,3]]
Output: 3
```

**Example 3:**
```
Input: n = 3, trust = [[1,3],[2,3],[3,1]]
Output: -1
```
## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/find-the-town-judge/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
