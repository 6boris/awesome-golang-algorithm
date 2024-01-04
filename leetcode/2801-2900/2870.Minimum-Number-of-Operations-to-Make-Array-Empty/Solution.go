package Solution

func Solution(nums []int) int {
	arr := make(map[int]int)
	for _, n := range nums {
		arr[n]++
	}

	ans := 0
	for _, n := range arr {
		if n == 1 {
			return -1
		}
		three := n / 3
		mod := n % 3
		if mod == 1 {
			three-- //挪出去一个3，组成3+1
			three += 2
		}
		if mod == 2 {
			three++
		}
		ans += three
	}
	return ans
}
