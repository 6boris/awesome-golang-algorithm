# [128. Longest Consecutive Sequence][title]

## Description
Given an unsorted array of integers, find the length of the longest consecutive elements sequence.

Your algorithm should run in O(n) complexity.

**Example:**

```
Input: [100, 4, 200, 1, 3, 2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
```

## 题意
> 给定一个未排序的整数数组，找出最长连续序列的长度。


## 题解

### 思路1
> 使用并查集，将相邻的元素合并，找到元素最多的集合。

```go
var p map[int]int
var rank map[int]int

func longestConsecutive(nums []int) int {
    l := len(nums)
    if l < 1 {
        return 0
    }
    p, rank = initialize(nums)
    hash := make(map[int]int)
    for _, v := range nums {
        hash[v] = 1
        if _, ok := hash[v-1]; ok {
            union(v, v-1)
        }
        if _, ok := hash[v+1]; ok {
            union(v, v+1)
        }
    }
    
    ans := 0
    for _, v := range rank {
        if v > ans {
            ans = v
        }
    }
    fmt.Println(rank)
    return ans
}

func initialize(nums []int) (map[int]int, map[int]int) {
    p, rank = make(map[int]int), make(map[int]int)
    for _, v := range nums {
        p[v]    = v
        rank[v] = 1
    }
    return p, rank
}

func find(x int) int {
    if p[x] != x {
        p[x] = find(p[x])
    }
    return p[x]
}

func union(x, y int) {
    x = find(x)
    y = find(y)
    if x != y {
        if rank[x] >= rank[y] {
            rank[x] += rank[y]
            p[y] = x
        } else {
            rank[y] += rank[x]
            p[x] = y
        }
    }
}
```

### 思路2
> 思路2
```go

```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/longest-consecutive-sequence/description/  
[me]: https://github.com/kylesliu/awesome-golang-algorithm
