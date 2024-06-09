package Solution

func Solution(nums []int) int {
	count := make([]int, 101)
	for _, n := range nums {
		count[n]++
	}

	ans := 0
	freq := -1
	for _, c := range count {
		if c == 0 {
			continue
		}
		if c > freq {
			ans = 0
			freq = c
		}
		if c == freq {
			ans += freq
		}
	}
	return ans
}
