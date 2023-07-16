# [1125.Smallest Sufficient Team][title]

## Description
In a project, you have a list of required skills `req_skills`, and a list of people. The i<sup>th</sup> person `people[i]` contains a list of skills that the person has.

Consider a sufficient team: a set of people such that for every required skill in `req_skills`, there is at least one person in the team who has that skill. We can represent these teams by the index of each person.

- For example, `team = [0, 1, 3]` represents the people with skills `people[0]`, `people[1]`, and `people[3]`.

Return any sufficient team of the smallest possible size, represented by the index of each person. You may return the answer in **any order**.

It is **guaranteed** an answer exists.

**Example 1:**

```
Input: req_skills = ["java","nodejs","reactjs"], people = [["java"],["nodejs"],["nodejs","reactjs"]]
Output: [0,2]
```

**Example 2:**

```
Input: req_skills = ["algorithms","math","java","reactjs","csharp","aws"], people = [["algorithms","math","java"],["algorithms","math","reactjs"],["java","csharp","aws"],["reactjs","csharp"],["csharp","math"],["aws","java"]]
Output: [1,2]
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/smallest-sufficient-team/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
