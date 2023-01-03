package Solution

func Solution(strs []string) int {
	ans := 0
	rows := len(strs)
	cols := len(strs[0])
	for col := 0; col < cols; col++ {
		c := strs[0][col]
		for row := 1; row < rows; row++ {
			if strs[row][col] < c {
				ans++
				break
			}
			c = strs[row][col]
		}
	}
	return ans
}
