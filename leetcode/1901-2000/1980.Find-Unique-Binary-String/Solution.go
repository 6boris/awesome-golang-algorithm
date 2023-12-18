package Solution

func addOne(bs []byte) {
	for i := len(bs) - 1; i >= 0; i-- {
		if bs[i] == '0' {
			bs[i] = '1'
			break
		}
		bs[i] = '0'
	}
}
func Solution(nums []string) string {
	l := len(nums)
	bs := make([]byte, l)
	set := make(map[string]struct{})
	for i := range bs {
		bs[i] = '0'
		set[nums[i]] = struct{}{}
	}

	for i := 0; i < (1<<l)-1; i++ {
		if _, ok := set[string(bs)]; !ok {
			break
		}
		addOne(bs)
	}
	return string(bs)
}
