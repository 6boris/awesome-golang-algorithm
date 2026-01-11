package Solution

const mod1573 = 1e9 + 7

func Solution(s string) int {
	ones := []int{}
	bs := []byte(s)
	for idx, b := range bs {
		if b == '1' {
			ones = append(ones, idx)
		}
	}
	if len(ones)%3 != 0 {
		return 0
	}
	var ret, l int
	if len(ones) == 0 {
		l = len(bs)
		for i := 0; i < l-2; i++ {
			left := l - i - 1
			ret = (ret + left - 1) % mod1573
		}
	} else {
		l = len(ones)
		spite := l / 3
		left := ones[spite] - ones[spite-1]
		right := ones[l-spite] - ones[l-spite-1]
		ret = (left * right) % mod1573
	}
	return ret
}
