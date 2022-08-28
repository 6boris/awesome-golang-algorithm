package Solution

func Solution(s string) int {
	pz, po := 0, 0
	if s[0] == '1' {
		pz = 1
	} else {
		po = 1
	}
	for idx := 1; idx < len(s); idx++ {
		if s[idx] == '1' {
			pre := po
			po = pz
			pz = pre + 1
			continue
		}
		pre := pz
		pz = po
		po = pre + 1
	}
	if pz > po {
		return po
	}
	return pz
}

func Solution2(s string) int {
	case1, b1 := 0, uint8(0) // 0101.....
	case2, b2 := 0, uint8(1) // 1010
	for _, is := range s {
		if uint8(is-'0') != b1 {
			case1++
		}
		b1 = 1 - b1
	}
	for _, is := range s {
		if uint8(is-'0') != b2 {
			case2++
		}
		b2 = 1 - b2
	}
	if case1 > case2 {
		return case2
	}
	return case1

}
