package Solution

const mod = 1000000007

func Solution(points [][]int) int {
	columns := make(map[int]int)
	for _, p := range points {
		columns[p[1]]++
	}

	pairs := []int{}
	sum, count := 0, 0
	for _, c := range columns {
		if c < 2 {
			continue
		}
		count = c * (c - 1) / 2
		sum = (sum + count) % mod
		pairs = append(pairs, count)
	}
	var ret int
	for _, pair := range pairs {
		ret = (ret + pair*((sum-pair+mod)%mod)) % mod
	}

	// 使用模逆元乘以 1/2
	inv2 := (mod + 1) / 2
	return (ret * inv2) % mod
}
