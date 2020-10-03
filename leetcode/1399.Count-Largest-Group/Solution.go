package Solution

func Solution(n int) int {
	hash := make(map[int]int)
	max := -1
	for i := 1; i <= n; i++ {
		var d int
		if i < 10 {
			d = i
		} else {
			t := i
			for t > 0 {
				d += t % 10
				t /= 10
			}
		}
		hash[d]++
		if hash[d] > max {
			max = hash[d]
		}
	}
	ans := 0
	for _, val := range hash {
		if val == max {
			ans++
		}
	}
	return ans
}
