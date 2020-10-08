package Solution

func singleNumber(nums []int) int {
	m := make(map[int]int, len(nums))
	for _, v := range nums {
		m[v]++
	}
	for v := range m {
		if m[v] == 1 {
			return v
		}
	}
	return 0
}

//	只循环一次
// 3 ∗ (a+b+c)−(a+a+b+b+c) = 2 * c
func singleNumber2(nums []int) int {
	m := make(map[int]int, len(nums))
	sum1, sum2 := 0, 0
	for _, v := range nums {
		if _, ok := m[v]; !ok {
			m[v]++
			sum1 += v
		}
		sum2 += v
	}
	return (3*sum1 - sum2) / 2
}
