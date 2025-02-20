package Solution

func Solution(arr []int, mat [][]int) int {
	m, n := len(mat), len(mat[0])
	rows, cols := make([]int, m), make([]int, n)
	key := make(map[int][2]int)
	for i := range m {
		for j := range n {
			key[mat[i][j]] = [2]int{i, j}
		}
	}
	for i := range arr {
		pos := key[arr[i]]
		rows[pos[0]]++
		cols[pos[1]]++
		if rows[pos[0]] == n || cols[pos[1]] == m {
			return i
		}
	}
	return len(arr) - 1
}
