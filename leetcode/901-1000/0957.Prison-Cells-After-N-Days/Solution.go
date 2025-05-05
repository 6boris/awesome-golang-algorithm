package Solution

func Solution(cells []int, n int) []int {
	start := [8]int{}
	for i := range cells {
		start[i] = cells[i]
	}
	days := [][8]int{start}
	day := 0
	cache := map[[8]int]int{
		start: day,
	}

	var convert func([8]int) [8]int
	convert = func(now [8]int) [8]int {
		var res [8]int
		res[0], res[7] = 0, 0
		for i := 1; i < 7; i++ {
			if now[i-1] == now[i+1] {
				res[i] = 1
			}
		}
		return res
	}
	var r [8]int
	var loopStart, loopLen int
	for {
		day++
		r = convert(start)
		if day == n {
			return r[:]
		}
		if d, ok := cache[r]; ok {
			loopStart = d
			loopLen = day - d
			break
		}
		cache[r] = day
		start = r
		days = append(days, r)
	}
	target := n - loopStart
	target %= loopLen
	return days[loopStart+target][:]
}
