package Soluation

func maxAreaOfIsland(grid [][]int) int {
    ans := 0
    height, width := len(grid), len(grid[0])

    for i := 0; i < height; i++ {
        for j := 0; j < width; j++ {
            tmp := 0;
            if 1 == grid[i][j] {
                tmp = dfs(grid, i, j)
                if tmp > ans {
                    ans = tmp
                }
            }
        }
    }
    return ans
}

func dfs(grid [][]int, i, j int) int {
    res := 0
    if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) && 1 == grid[i][j] {
        grid[i][j] = 0
        up := dfs(grid, i-1, j)
        down := dfs(grid, i+1, j)
        left := dfs(grid, i, j-1)
        right := dfs(grid, i, j+1)
        res = up + down + left + right + 1
    }
    return res
}