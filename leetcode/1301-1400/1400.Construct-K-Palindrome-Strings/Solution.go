package Solution

func Solution(s string, k int) bool {
	if len(s) < k {
		return false
	}
	count := [26]int{}
	for _, b := range s {
		count[b-'a']++
	}
	odd, even := 0, 0
	for _, n := range count {
		if n == 0 {
			continue
		}
		even += n
		if n&1 == 1 {
			odd++
		}
	}
	if odd > k {
		return false
	}
	if k == odd {
		return true
	}
	return k <= even
}
