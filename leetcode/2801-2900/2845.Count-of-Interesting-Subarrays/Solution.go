package Solution

func Solution(nums []int, modulo int, k int) int64 {
	cache := make(map[int]int)
	var res int64 = 0
	cnt := 0
	cache[0] = 1
	for _, n := range nums {
		if n%modulo == k {
			cnt++
		}
		res += int64(cache[(cnt-k+modulo)%modulo])
		cache[cnt%modulo]++
	}
	return res
}
