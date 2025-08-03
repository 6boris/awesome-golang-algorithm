package Solution

func pow(x int) int {
	base := 1

	for i := 1; i <= x; i++ {
		base *= 2
	}
	return base
}

func Solution(s string, n int) bool {
	count := make(map[int]struct{})
	ls := len(s)

	loop := min(31, ls)
	for i := 1; i <= loop; i++ {
		base := 0
		for j := 0; j < i; j++ {
			base = base*2 + int(s[j]-'0')
		}
		if base != 0 && base <= n {
			count[base] = struct{}{}
		}
		b2 := pow(i - 1)
		start, end := 0, i
		for ; end < ls; start, end = start+1, end+1 {
			if s[start] == '1' {
				base -= b2
			}
			now := base << 1
			now += int(s[end] - '0')
			base = now
			if now == 0 || now > n {
				continue
			}
			count[now] = struct{}{}
		}
	}
	return len(count) == n
}
