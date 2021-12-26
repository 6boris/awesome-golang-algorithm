package Solution

func Solution(points [][]int, k int) [][]int {
	result := make([][]int, 0)
	for {
		m := kClosestPartion(points)
		if m == k-1 {
			result = append(result, points[:m+1]...)
			return result
		}
		if m < k {
			result = append(result, points[:m+1]...)
			points = points[m+1:]
			k -= m + 1
			continue
		}
		points = points[:m]
	}
}

func kClosestPartion(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	x, y := points[0][0], points[0][1]
	cmpObj, idx := x*x+y*y, 0

	for walker := 1; walker < len(points); walker++ {
		x, y = points[walker][0], points[walker][1]
		dis := x*x + y*y
		if dis < cmpObj {
			idx++
			points[idx], points[walker] = points[walker], points[idx]
		}
	}
	points[idx], points[0] = points[0], points[idx]
	return idx
}
