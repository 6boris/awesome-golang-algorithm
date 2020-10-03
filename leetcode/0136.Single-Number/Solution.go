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
// 2∗(a+b+c)−(a+a+b+b+c) = c
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
	return 2*sum1 - sum2
}

//	用异或
//　0^0 = 0，
//　1^0 = 1，
//　0^1 = 1，
//　1^1 = 0

func singleNumber3(nums []int) int {
	ans := 0
	for _, v := range nums {
		ans ^= v
	}
	return ans
}
