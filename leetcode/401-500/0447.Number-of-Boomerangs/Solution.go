package Solution

func Solution(points [][]int) int {
	d := make(map[int][][2]int)
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			x := points[i][0] - points[j][0]
			y := points[i][1] - points[j][1]
			dis := x*x + y*y
			if _, ok := d[dis]; !ok {
				d[dis] = make([][2]int, 0)
			}
			d[dis] = append(d[dis], [2]int{i, j})
		}
	}
	ans := 0
	for _, pos := range d {
		l := len(pos)
		for i := 0; i < l-1; i++ {
			for j := i + 1; j < l; j++ {
				if pos[i][0] == pos[j][0] || pos[i][1] == pos[j][1] || pos[i][1] == pos[j][0] {
					ans += 2
				}
			}
		}
	}
	return ans
}
