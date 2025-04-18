package Solution

func Solution(count []int) []float64 {
	var _max, _min, _mode, _modeCount int
	_min = -1
	total := int64(0)
	sum := int64(0)

	for i, c := range count {
		if c == 0 {
			continue
		}
		sum += int64(i) * int64(c)
		total += int64(c)

		_max = max(_max, i)
		if _min == -1 || _min > i {
			_min = i
		}
		if c > _modeCount {
			_modeCount = c
			_mode = i
		}
	}
	_avg := float64(sum) / float64(total)
	var a float64
	start := total / 2
	if total&1 == 0 {
		start--
	}
	used := int64(0)
	justSelected := false
	for i, c := range count {
		if c == 0 {
			continue
		}
		if justSelected {
			a += float64(i)
			a /= 2
			break
		}
		next := used + int64(c)
		if next-1 < start {
			used = next
			continue
		}

		diff := next - start
		a += float64(i)
		if total&1 == 1 {
			break
		}
		if diff >= 2 {
			a += a
			a /= 2
			break
		}
		justSelected = true
	}

	return []float64{float64(_min), float64(_max), _avg, a, float64(_mode)}
}
