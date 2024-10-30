package Solution

func Solution(n int, k int) byte {
	alloc := 1
	for i := 2; i <= n; i++ {
		alloc = (alloc * 2) + 1
	}
	bs := make([]byte, alloc)
	bs[0] = '0'
	index := 1
	for i := 2; i <= n; i++ {
		bs[index] = '1'
		end := index
		index++
		for j := end - 1; j >= 0; j-- {
			bs[index] = '0'
			if bs[j] == '0' {
				bs[index] = '1'
			}
			index++
		}
	}
	return bs[k-1]
}
