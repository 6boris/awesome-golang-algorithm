package Solution

func Solution(nums []int, original int) int {
	var (
		ret int  = original
		ok  bool = false
		end int  = 0
	)
	m := make(map[int]struct{})
	for _, n := range nums {
		end = max(end, n)
		m[n] = struct{}{}
	}
	for ; ret <= end; ret *= 2 {
		_, ok = m[ret]
		if !ok {
			break
		}
	}
	return ret
}
