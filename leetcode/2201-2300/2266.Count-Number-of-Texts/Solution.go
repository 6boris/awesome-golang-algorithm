package Solution

const mod2266 = 1e9 + 7

type fn2266 func(int) int

func for3(n int) int {
	if n == 0 {
		return 1
	}
	if n <= 2 {
		return n
	}
	a, b, c := 1, 1, 2
	tmp := 0
	for i := 3; i <= n; i++ {
		tmp = (a + b + c) % mod2266
		a, b, c = b, c, tmp
	}
	return c
}

func for4(n int) int {
	if n == 0 {
		return 1
	}
	if n <= 2 {
		return n
	}
	if n == 3 {
		return 4
	}
	a, b, c, d := 1, 1, 2, 4
	tmp := 0
	for i := 4; i <= n; i++ {
		tmp = (a + b + c + d) % mod2266
		a, b, c, d = b, c, d, tmp
	}
	return d
}

func Solution(pressedKeys string) int {
	bs := []byte(pressedKeys)
	prev := bs[0]

	ret, cnt := 1, 1
	cache := make(map[[2]int]int)
	var which int
	fn := []fn2266{for4, for3}

	for i := 1; i < len(bs); i++ {
		if bs[i] == prev {
			cnt++
			continue
		}

		which = 3
		if prev == '7' || prev == '9' {
			which = 4
		}
		key := [2]int{which, cnt}

		v, ok := cache[key]
		if !ok {
			v = fn[which&1](cnt)
			cache[key] = v
		}
		ret = (ret * v) % mod2266

		cnt = 1
		prev = bs[i]
	}
	which = 3
	if prev == '7' || prev == '9' {
		which = 4
	}
	key := [2]int{which, cnt}

	v, ok := cache[key]
	if !ok {
		v = fn[which&1](cnt)
	}
	ret = (ret * v) % mod2266

	return ret
}
