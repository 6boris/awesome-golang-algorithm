package Solution

func Solution(nums []int) []int {
	cnt := make(map[int]int)
	for _, n := range nums {
		cnt[n]++
	}
	ans := make([]int, 0)
	for _, n := range nums {
		if cnt[n] != 1 {
			continue
		}
		_, ok1 := cnt[n-1]
		_, ok2 := cnt[n+1]
		if !ok1 && !ok2 {
			ans = append(ans, n)
		}
	}
	return ans
}
