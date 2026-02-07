package Solution

func Solution(target string) []string {
	var (
		ret []string
		cur string
	)
	index := 0
	letters := [26]string{}
	chars := [26]byte{}
	for i := range 26 {
		letters[i] = string([]byte{byte('a' + i)})
		chars[i] = byte('a' + i)
	}
	bs := []byte(target)
	for i := range bs {
		for ; chars[index] != bs[i]; index = (index + 1) % 26 {
			ret = append(ret, cur+letters[index])
		}
		cur += letters[bs[i]-'a']
		ret = append(ret, cur)
		index = 0
	}
	return ret
}
