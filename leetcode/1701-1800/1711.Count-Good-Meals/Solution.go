package Solution

const mod = 1000000007

func Solution(deliciousness []int) int {
	cur := int64(1)
	cnt := make(map[int64]int)
	for _, n := range deliciousness {
		cnt[int64(n)]++
	}
	ans := 0
	for i := 0; i <= 41; i++ {
		skip := map[int64]struct{}{}
		for d, n := range cnt {
			if _, ok := skip[d]; ok {
				continue
			}
			diff := cur - d
			if v, ok := cnt[diff]; ok {
				skip[diff] = struct{}{}
				if diff == d {
					ans = (ans + n*(n-1)/2) % mod
				} else {
					ans = (ans + n*v) % mod
				}
			}
		}
		cur <<= 1
	}
	return ans
}
