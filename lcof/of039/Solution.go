package Solution

func majorityElement(nums []int) int {
	cur, count := 0, 0
	for _, v := range nums {
		if count == 0 {
			cur = v
		}
		if v == cur {
			count++
		} else {
			count--
		}
	}
	return cur
}
