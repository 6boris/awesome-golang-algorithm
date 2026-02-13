package Solution

type fn func(int) bool

func Solution(arr []int) int {
	cnt, l := 1, len(arr)

	var (
		ret int
		gt  = func(i int) bool {
			return arr[i] > arr[i+1]
		}
		lt = func(i int) bool {
			return arr[i] < arr[i+1]
		}
	)

	funcs := []fn{lt, gt}
	for i := 0; i < l-1; i++ {
		if funcs[i&1](i) {
			cnt++
			continue
		}
		ret = max(ret, cnt)
		cnt = 1
	}
	ret = max(ret, cnt)

	cnt = 1
	for i := 0; i < l-1; i++ {
		if funcs[(i+1)&1](i) {
			cnt++
			continue
		}
		ret = max(ret, cnt)
		cnt = 1
	}
	ret = max(ret, cnt)
	return ret
}
