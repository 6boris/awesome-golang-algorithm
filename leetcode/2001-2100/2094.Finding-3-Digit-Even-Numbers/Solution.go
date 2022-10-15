package Solution

func Solution(digits []int) []int {

	bucket := make([]int, 10)

	for _, i := range digits {
		bucket[i]++
	}

	ans := make([]int, 0)
	for idx := 100; idx < 1000; idx += 2 {
		a := idx / 100
		b := idx % 100
		c := b / 10
		d := b % 10

		bucket[a]--
		bucket[c]--
		bucket[d]--
		if bucket[a] >= 0 && bucket[c] >= 0 && bucket[d] >= 0 {
			ans = append(ans, idx)
		}
		bucket[a]++
		bucket[c]++
		bucket[d]++
	}
	return ans
}
