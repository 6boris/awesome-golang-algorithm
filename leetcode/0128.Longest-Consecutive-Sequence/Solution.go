package Soluation

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