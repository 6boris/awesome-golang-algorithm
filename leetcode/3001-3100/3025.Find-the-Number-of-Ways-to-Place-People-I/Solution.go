package Solution

func Solution(points [][]int) int {
	ans := 0
	l := len(points)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			// 检查他俩是否ok
			if i == j || !(points[i][0] <= points[j][0] && points[i][1] >= points[j][1]) {
				// 不在左上角
				continue
			}
			k := 0
			for ; k < l; k++ {
				if k == i || k == j {
					continue
				}
				x := points[k][0] <= points[j][0] && points[k][0] >= points[i][0]
				y := points[k][1] >= points[j][1] && points[k][1] <= points[i][1]
				if x && y {
					break
				}
			}
			if k == l {
				ans++
			}

		}
	}
	return ans
}
