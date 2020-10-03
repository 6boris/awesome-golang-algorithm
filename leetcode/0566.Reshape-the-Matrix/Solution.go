package Solution

func Solution(nums [][]int, r int, c int) [][]int {
	m, n := len(nums), len(nums[0])
	if m*n != r*c {
		return nums
	}
	newNums := make([][]int, r)
	for i := range newNums {
		newNums[i] = make([]int, c)
	}
	k := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			r_, c_ := k/c, k%c
			newNums[r_][c_] = nums[i][j]
			k++
		}
	}
	return newNums
}
