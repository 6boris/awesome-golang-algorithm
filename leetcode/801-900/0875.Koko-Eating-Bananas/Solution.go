package Solution

func Solution(piles []int, h int) int {
	minSpeed := 1
	maxSpeed := -1
	for _, t := range piles {
		if t > maxSpeed {
			maxSpeed = t
		}
	}

	ans := 1
	for minSpeed <= maxSpeed {
		midSpeed := (maxSpeed-minSpeed)/2 + minSpeed
		cost := 0
		for _, t := range piles {
			a := t / midSpeed
			if t%midSpeed != 0 {
				a++
			}
			cost += a
			if cost > h {
				break
			}
		}
		if cost <= h {
			ans = midSpeed
			maxSpeed = midSpeed - 1
		} else {
			minSpeed = midSpeed + 1
		}
	}
	return ans
}
