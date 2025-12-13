package Solution

func Solution(nums []int) int64 {
	l, cnt := len(nums), 0
	ret := int64(l)
	start, end := 0, 1

	for ; end < l; end++ {
		if nums[end] != nums[end-1] {
			continue
		}

		length := end - start
		cnt = (length - 1) * length / 2
		ret += int64(cnt)

		start = end
	}

	length := end - start
	cnt = (length - 1) * length / 2
	ret += int64(cnt)

	return ret
}
