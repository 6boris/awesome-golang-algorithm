package Solution

func Solution(strs []string) int {
	var ret, row int
	rows := len(strs)
	cols := len(strs[0])

	sorted := make([]bool, rows-1)
	/*
		bbjwefkpb
		axmksfchw
	*/
	for col := 0; col < cols; col++ {
		for row = 0; row < rows-1; row++ {
			if !sorted[row] && strs[row][col] > strs[row+1][col] {
				ret++
				break
			}
		}
		if row == rows-1 {
			for row = 0; row < rows-1; row++ {
				if !sorted[row] && strs[row][col] < strs[row+1][col] {
					sorted[row] = true
				}
			}
		}

	}
	return ret
}
