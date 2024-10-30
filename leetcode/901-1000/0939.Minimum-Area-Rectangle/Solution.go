package Solution

func Solution(points [][]int) int {
	keys := make(map[[2]int]struct{})
	for _, xy := range points {
		keys[[2]int{xy[0], xy[1]}] = struct{}{}
	}

	ans := 0
	for i, xy := range points {
		for j, other := range points {
			if i == j {
				continue
			}
			if xy[0] == other[0] || xy[1] == other[1] {
				continue
			}
			x1, y1, x2, y2 := xy[0], other[1], other[0], xy[1]
			_, ok1 := keys[[2]int{x1, y1}]
			if ok1 {
				if _, ok2 := keys[[2]int{x2, y2}]; ok2 {
					width := x1 - x2
					if width < 0 {
						width = -width
					}
					height := y1 - y2
					if height < 0 {
						height = -height
					}
					area := width * height
					if ans == 0 || ans > area {
						ans = area
					}
				}
			}
		}
	}
	return ans
}
