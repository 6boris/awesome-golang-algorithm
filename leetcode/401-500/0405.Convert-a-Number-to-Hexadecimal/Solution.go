package Solution

func Solution(num int) string {
	n2b := map[int]byte{
		0: '0', 1: '1', 2: '2', 3: '3', 4: '4', 5: '5', 6: '6', 7: '7',
		8: '8', 9: '9', 10: 'a', 11: 'b', 12: 'c', 13: 'd', 14: 'e', 15: 'f',
	}
	if num >= 0 {
		bs := make([]byte, 0)
		for num >= 16 {
			mod := num % 16
			num /= 16
			bs = append(bs, n2b[mod])
		}
		bs = append(bs, n2b[num])
		for s, e := 0, len(bs)-1; s < e; s, e = s+1, e-1 {
			bs[s], bs[e] = bs[e], bs[s]
		}
		return string(bs)
	}
	num = -num
	bs := make([]uint8, 32)
	bs[0] = 1
	end := 31
	for num >= 2 {
		mod := num % 2
		bs[end] = uint8(mod)
		num /= 2
		end--
	}
	bs[end] = uint8(num)
	cf := uint8(1)
	for idx := 31; idx > 0; idx-- {
		bs[idx] ^= uint8(1)
		bs[idx] += cf
		cf = bs[idx] / 2
		bs[idx] %= 2
	}
	power := [4]uint8{1, 2, 4, 8}
	result := make([]byte, 8)
	bsIdx, idx := 7, 31
	for idx > 0 {
		base := uint8(0)
		for i := 0; i < 4; i++ {
			base += bs[idx-i] * power[i]
		}
		result[bsIdx] = n2b[int(base)]
		bsIdx--
		idx -= 4
	}
	// ^=1
	return string(result)
}

func Solution2(num int) string {
	maxInt := 2147483647  // 2**31-1
	minInt := -maxInt - 1 // -2**31
	compare := 4294967295 // 2**32-1
	if num > maxInt || num < minInt {
		return ""
	}
	if num == 0 {
		return "0"
	}
	if num < 0 {
		num = num + 1 + compare
	}
	numMap := map[int]byte{
		0: '0', 1: '1', 2: '2', 3: '3', 4: '4',
		5: '5', 6: '6', 7: '7', 8: '8', 9: '9',
		10: 'a', 11: 'b', 12: 'c', 13: 'd', 14: 'e',
		15: 'f',
	}
	res := make([]byte, 0)
	for num > 0 {
		remainer := num % 16
		num = num / 16
		res = append(res, numMap[remainer])
	}
	for s, e := 0, len(res)-1; s < e; s, e = s+1, e-1 {
		res[s], res[e] = res[e], res[s]
	}
	return string(res)
}
