package Solution

func Solution(n int) bool {
	arr := make([]int, 15)
	arr[0] = 1
	for i := 1; i < 15; i++ {
		arr[i] = arr[i-1] * 3
	}
	var ok func(int, int) bool
	ok = func(index, cur int) bool {
		if cur == 0 {
			return true
		}
		if index == 15 {
			return false
		}

		return ok(index+1, cur-arr[index]) || ok(index+1, cur)
	}
	return ok(0, n)
}
