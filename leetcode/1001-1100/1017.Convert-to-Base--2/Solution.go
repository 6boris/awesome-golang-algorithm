package Solution

func Solution(n int) string {
	if n == 0 {
		return "0"
	}
	var result []byte
	for n != 0 {
		remainder := n % -2
		n = n / -2

		if remainder < 0 {
			remainder += 2
			n += 1
		}
		result = append(result, byte(remainder+48))
	}
	for s, e := 0, len(result)-1; s < e; s, e = s+1, e-1 {
		result[s], result[e] = result[e], result[s]
	}
	return string(result)

}
