# [547. Friend Circles][title]

## Description

There are N students in a class. Some of them are friends, while some are not. Their friendship is transitive in nature. For example, if A is a direct friend of B, and B is a direct friend of C, then A is an indirect friend of C. And we defined a friend circle is a group of students who are direct or indirect friends.

Given a N*N matrix M representing the friend relationship between students in the class. If M[i][j] = 1, then the ith and jth students are direct friends with each other, otherwise not. And you have to output the total number of friend circles among all the students.

**Example 1:**
```
Input: 
[[1,1,0],
 [1,1,0],
 [0,0,1]]
Output: 2
Explanation:The 0th and 1st students are direct friends, so they are in a friend circle. 
The 2nd student himself is in a friend circle. So return 2.
```

**Example 2:**
```
Input: 
[[1,1,0],
 [1,1,1],
 [0,1,1]]
Output: 1
Explanation:The 0th and 1st students are direct friends, the 1st and 2nd students are direct friends, 
so the 0th and 2nd students are indirect friends. All of them are in the same friend circle, so return 1.
```

## 题意
> 判断几个学生是否在同一个圈子。

## 题解

### 思路 1
> 使用并查集

```go
var p []int
var r []int
var ans int

func findCircleNum(M [][]int) int {
    ans = len(M)
    initialize(ans)

    for i := 0; i < len(M); i++ {
        for j := 0; j < i; j++ {
            if 1 == M[i][j] {
                union(i, j)
            }
        }
    }
    return ans
}

func initialize(l int) {
    p = make([]int, l)
    r = make([]int, l)
    for i, _ := range p {
        p[i] = i
        r[i] = 1
    }
    fmt.Println(p, r)
}

func find(x int) int {
    if x != p[x] {
        p[x] = find(p[x])
    }
    return p[x]
}

func union(x, y int) {
    x, y = find(x), find(y)
    if x != y {
        if r[x] <= r[y] {
            p[x] = y
            r[y] += r[x]
        } else {
            p[y] = x
            r[x] += r[y]
        }
        ans--
    }
}
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-leetcode][me]

[title]: https://leetcode.com/problems/friend-circles/description/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
