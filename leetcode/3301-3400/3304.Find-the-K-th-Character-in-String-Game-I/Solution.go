package Solution

func Solution(k int) byte {
	str := "a"
	for len(str) < k {
		tmp := []byte(str)
		for i := range tmp {
			tmp[i] = (tmp[i]-'a'+1)%26 + 'a'
		}
		str += string(tmp)
	}
	return str[k-1]
}
