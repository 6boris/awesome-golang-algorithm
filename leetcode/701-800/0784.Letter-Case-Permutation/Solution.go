package Solution

func Solution(s string) []string {
	bs := []byte(s)
	ans := make([]string, 0)
	help(bs, 0, &ans)
	return ans
}

func help(bs []byte, index int, ans *[]string) {
	if index == len(bs) {
		*ans = append(*ans, string(bs))
		return
	}
	// diff = 32
	if bs[index] >= '0' && bs[index] <= '9' {
		help(bs, index+1, ans)
		return
	}
	source := bs[index]
	help(bs, index+1, ans)

	if bs[index] >= 'a' && bs[index] <= 'z' {
		bs[index] = bs[index] - 32
	} else if bs[index] >= 'A' && bs[index] <= 'Z' {
		bs[index] = bs[index] + 32
	}
	help(bs, index+1, ans)
	bs[index] = source

}
