package Solution

func Solution(s string) int {
	ans := 0
	shift := map[byte]int{
		'a': 0, 'e': 1, 'i': 2, 'o': 3, 'u': 4,
	}
	b := 0
	count := map[int]int{0: -1}
	for i, c := range s {
		if pos, ok := shift[byte(c)]; ok {
			b ^= 1 << pos
			if _, ok := count[b]; !ok {
				count[b] = i
			}
		}
		if r := i - count[b]; r > ans {
			ans = r
		}
	}
	return ans
}
