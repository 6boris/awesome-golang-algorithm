package Solution

func toKey(n int) [10]int {
	tmp := [10]int{}
	for n > 0 {
		cur := n % 10
		tmp[cur]++
		n /= 10
	}
	return tmp
}
func Solution(n int) bool {
	if n&(n-1) == 0 {
		return true
	}

	cache := map[[10]int]struct{}{}
	key := toKey(n)
	end := 0
	for i := 9; i >= 0; i-- {
		count := key[i]
		for ; count > 0; count-- {
			end = end*10 + i
		}
	}
	shift := 0
	cur := 1
	for cur <= end {
		t := toKey(cur)
		cache[t] = struct{}{}
		shift++
		cur = 1 << shift
	}
	_, ok := cache[toKey(n)]
	return ok
}
