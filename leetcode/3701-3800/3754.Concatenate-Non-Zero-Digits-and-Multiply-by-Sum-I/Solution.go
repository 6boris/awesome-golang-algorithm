package Solution

func Solution(n int) int64 {
	var (
		base        int64 = 1
		sum, x, tmp int64
	)
	i64n := int64(n)
	for ; i64n > 0; i64n /= 10 {
		tmp = i64n % 10
		if tmp == 0 {
			continue
		}

		sum += tmp
		x = base*int64(tmp) + x
		base *= 10
	}
	return sum * x
}
