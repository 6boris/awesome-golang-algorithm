package Solution

func Solution(digits []int) int {
	cnt := [10]int{}
	for _, d := range digits {
		cnt[d]++
	}

	var ret int
	for num := 100; num < 1000; num += 2 {
		a := num / 100
		b := (num / 10) % 10
		c := num % 10

		temp := cnt
		temp[a]--
		temp[b]--
		temp[c]--

		if temp[a] >= 0 && temp[b] >= 0 && temp[c] >= 0 {
			ret++
		}
	}
	return ret
}
