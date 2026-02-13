package Solution

func Solution(bottomLeft [][]int, topRight [][]int) int64 {
	n := len(bottomLeft)
	maxSide := int64(0)
	var (
		interLeft, interRight, interBottom, interTop int64
	)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// 计算两个矩形的交集边界
			interLeft = int64(max(bottomLeft[i][0], bottomLeft[j][0]))
			interRight = int64(min(topRight[i][0], topRight[j][0]))
			interBottom = int64(max(bottomLeft[i][1], bottomLeft[j][1]))
			interTop = int64(min(topRight[i][1], topRight[j][1]))

			// 计算交集宽高
			interWidth := interRight - interLeft
			interHeight := interTop - interBottom

			if interWidth > 0 && interHeight > 0 {
				side := min(interWidth, interHeight)
				maxSide = max(maxSide, side)
			}
		}
	}
	return maxSide * maxSide
}
