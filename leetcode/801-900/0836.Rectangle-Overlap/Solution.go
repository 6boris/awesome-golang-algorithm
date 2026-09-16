package Solution

func Solution(rec1 []int, rec2 []int) bool {
	return !(rec1[2] <= rec2[0] || // rec1 在 rec2 的左侧
		rec1[0] >= rec2[2] || // rec1 在 rec2 的右侧
		rec1[3] <= rec2[1] || // rec1 在 rec2 的下方
		rec1[1] >= rec2[3]) // rec1 在 rec2 的上方
}
