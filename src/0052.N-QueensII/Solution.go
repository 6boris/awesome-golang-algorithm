package Solution

func totalNQueens(n int) int {
	count := 0
	cols := make([]bool, n)
	for i := 0; i < len(cols); i++ {
		cols[i] = false
	}

	d1 := make([]bool, 2*n)
	d2 := make([]bool, 2*n)
	for i := 0; i < len(cols); i++ {
		d1[i] = false
		d2[i] = false
	}

	helper(0, cols, d1, d2, n, &count)

	return count
}

func helper(row int, cols []bool, d1 []bool, d2 []bool, n int, count *int) {
	if row == n {
		*count++
	}
	for col := 0; col < n; col++ {
		id1 := col - row + n
		id2 := col + row
		if cols[col] || d1[id1] || d2[id2] {
			continue
		}
		cols[col] = true
		d1[id1] = true
		d2[id2] = true
		helper(row+1, cols, d1, d2, n, count)
		cols[col] = false
		d1[id1] = false
		d2[id2] = false
	}
}
