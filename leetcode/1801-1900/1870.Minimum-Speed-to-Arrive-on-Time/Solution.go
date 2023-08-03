package Solution

func Solution(dist []int, hour float64) int {
	l, r := 1, 1000000001
	var canReach func(int) bool
	canReach = func(speed int) bool {
		cost := float64(0)
		for idx, dis := range dist {
			tmp := float64(dis) / float64(speed)
			if dis%speed != 0 && idx != len(dist)-1 {
				tmp = float64(dis/speed + 1)
			}
			cost += tmp
		}
		return hour >= cost
	}

	ans := -1
	for l <= r {
		if l == r {
			if canReach(l) {
				ans = l
			}
			break
		}

		mid := r - (r-l)/2
		if canReach(mid) {
			r = mid - 1
			ans = mid
			continue
		}
		l = mid
	}
	return ans
}
