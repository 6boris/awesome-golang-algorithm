package Solution

func Solution(arr []int, start int) bool {
	l := len(arr)
	used := make([]bool, l)
	used[start] = true
	queue := []int{start}
	for len(queue) > 0 {
		nq := make([]int, 0)
		for _, q := range queue {
			if arr[q] == 0 {
				return true
			}
			if left := q - arr[q]; left >= 0 && !used[left] {
				nq = append(nq, left)
				used[left] = true
			}
			if right := q + arr[q]; right < l && !used[right] {
				nq = append(nq, right)
				used[right] = true
			}
		}
		queue = nq
	}
	return false
}
