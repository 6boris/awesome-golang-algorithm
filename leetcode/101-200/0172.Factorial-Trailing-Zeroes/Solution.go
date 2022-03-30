package Solution

func Solution(n int) int {
	count := 0
	countOfFive := make(map[int]int)
	for step := 5; step <= n; step += 5 {
		countOfFive[step] = 1
		if c, ok := countOfFive[step/5]; ok {
			countOfFive[step] += c
		}
		count += countOfFive[step]
	}
	return count
}

func Solution2(n int) int {
	res := 0
	for n > 0 {
		res += n / 5
		n /= 5
	}
	return res
}
