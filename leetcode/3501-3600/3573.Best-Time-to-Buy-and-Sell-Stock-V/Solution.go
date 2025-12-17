package Solution

func Solution(prices []int, k int) int64 {
    n := len(prices)
    memo := make([][][]int64, n)
    for i := range memo {
        memo[i] = make([][]int64, k+1)
        for j := range memo[i] {
            memo[i][j] = make([]int64, 3)
            for s := range memo[i][j] {
                memo[i][j][s] = -1
            }
        }
    }
    
    var dfs func(int, int, int) int64
    dfs = func(i, j, state int) int64 {
        if j == 0 {
            return 0
        }
        if i == 0 {
            if state == 0 {
                return 0
            } else if state == 1 {
                return -int64(prices[0])
            } else {
                return int64(prices[0])
            }
        }
        if memo[i][j][state] != -1 {
            return memo[i][j][state]
        }
        
        p := prices[i]
        var res int64
        if state == 0 {
            res = max(dfs(i - 1, j, 0), max(dfs(i - 1, j, 1) + int64(p), dfs(i - 1, j, 2) - int64(p)))
        } else if state == 1 {
            res = max(dfs(i - 1, j, 1), dfs(i - 1, j - 1, 0) - int64(p))
        } else {
            res = max(dfs(i - 1, j, 2), dfs(i - 1, j - 1, 0) + int64(p))
        }
        memo[i][j][state] = res
        return res
    }
    
    return dfs(n-1, k, 0)
}
