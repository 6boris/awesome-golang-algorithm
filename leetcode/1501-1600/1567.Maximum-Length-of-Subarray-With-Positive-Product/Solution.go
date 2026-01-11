package Solution

func Solution(nums []int) int {
	var ret int
	cnt := 0
	start := -1
	first, last := -1, -1
	l := len(nums)
	for idx, n := range nums {
		if n == 0 {
			tmp := idx - start - 1
			if cnt&1 == 0 {
				ret = max(ret, tmp)
			} else {
				left := first - start
				right := idx - last
				ret = max(ret, tmp-left, tmp-right)
			}
			start = idx
			first, last, cnt = -1, -1, 0
			continue
		}
		if n > 0 {
			continue
		}
		if first == -1 {
			first = idx
		}
		last = idx
		cnt++
	}
	tmp := l - start - 1
	if cnt&1 == 0 {
		ret = max(ret, tmp)
	} else {
		left := first - start
		right := l - last
		ret = max(ret, tmp-left, tmp-right)
	}
	return ret
}
