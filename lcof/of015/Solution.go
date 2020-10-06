package Solution

func hammingWeight(num uint32) int {
	count := uint32(0)
	for num != 0 {
		count += num & 1

		num = num >> 1
	}
	return int(count)
}
