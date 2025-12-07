package Solution

func Solution(directions string) int {
	var ret int
	l := len(directions)
	points := [][2]int{}
	for i := 0; i < l-1; i++ {
		if directions[i] == 'R' && directions[i+1] == 'L' {
			points = append(points, [2]int{i, i + 1})
			continue
		}
		if directions[i] == 'S' {
			points = append(points, [2]int{i, l})
		}
	}
	if directions[l-1] == 'S' {
		points = append(points, [2]int{l - 1, l})
	}

	for _, p := range points {
		left := p[0] - 1
		right := p[0] + 1
		if p[1] != l {
			ret += 2
			right = p[1] + 1
		}
		for ; left >= 0 && directions[left] == 'R'; left-- {
			ret++
		}
		for ; right < l && directions[right] == 'L'; right++ {
			ret++
		}
	}
	return ret
}
