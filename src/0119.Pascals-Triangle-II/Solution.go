package Solution

func Solution(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	pt := make([]int, 0)
	for i := 0; i < rowIndex; i++ {
		tmp := make([]int, 0)
		tmp = append(tmp, 1)
		for j := 1; j < i+1; j++ {
			tmp = append(tmp, pt[j]+pt[j-1])
		}
		tmp = append(tmp, 1)
		pt = tmp
	}
	return pt
}
