package Solution

func Solution(nums []int) int {
	freq := make(map[int]int)
	for _, n := range nums {
		freq[n]++
	}

	values := make([]int, 0)
	for v := range freq {
		values = append(values, v)
	}
	var ret int
	pairs := make(map[int]int)
	for _, a := range values {
		for _, b := range values {
			pairs[a&b] += freq[a] * freq[b]
		}
	}
	for v, c := range pairs {
		for _, n := range values {
			if v&n == 0 {
				ret += c * freq[n]
			}
		}
	}
	return ret
}
