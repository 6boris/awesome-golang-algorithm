package Solution

func Solution(s string) bool {
	count := len(s) - 2
	bs := []byte(s)
	for i := count; i > 0; i-- {
		for j := 0; j <= i; j++ {
			bs[j] = byte((int(bs[j]-'0')+int(bs[j+1]-'0'))%10) + '0'
		}
	}
	return bs[0] == bs[1]
}
