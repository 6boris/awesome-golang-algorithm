package Solution

func Solution(name, typed string) bool {
	if name == typed {
		return true
	}

	nowByte := name[0]
	repeatBytes := 1
	typedIdx := 0
	typedRepeatBytes := 0
	ans := false
	for nameIdx := 1; nameIdx < len(name); nameIdx++ {
		if name[nameIdx] == nowByte {
			repeatBytes++
			continue
		}

		for ; typedIdx < len(typed) && typed[typedIdx] == nowByte; typedIdx++ {
			typedRepeatBytes++
		}

		if typedRepeatBytes < repeatBytes {
			return false
		}

		if repeatBytes < typedRepeatBytes {
			ans = true
		}

		typedRepeatBytes, repeatBytes = 0, 1
		nowByte = name[nameIdx]
	}
	for ; typedIdx < len(typed) && typed[typedIdx] == nowByte; typedIdx++ {
		typedRepeatBytes++
	}
	if typedRepeatBytes < repeatBytes || typedIdx != len(typed) {
		return false
	}

	if repeatBytes < typedRepeatBytes {
		ans = true
	}
	return ans
}

type byteRepeat struct {
	b byte
	c int
}

func checkPressTimes(s string) []byteRepeat {
	r := make([]byteRepeat, 0)
	if len(s) == 0 {
		return r
	}
	nowByte := s[0]
	repeatTimes := 1
	for idx := 1; idx < len(s); idx++ {
		if s[idx] == nowByte {
			repeatTimes++
			continue
		}

		r = append(r, byteRepeat{b: nowByte, c: repeatTimes})

		nowByte = s[idx]
		repeatTimes = 1
	}
	r = append(r, byteRepeat{b: nowByte, c: repeatTimes})
	return r
}

func Solution2(name string, typed string) bool {
	if name == typed {
		return true
	}
	baseName := checkPressTimes(name)
	baseTyped := checkPressTimes(typed)
	if len(baseTyped) != len(baseName) {
		return false
	}

	multiByte := false
	for idx := 0; idx < len(baseName); idx++ {
		if baseName[idx].b == baseTyped[idx].b {
			if baseName[idx].c > baseTyped[idx].c {
				return false
			}
			if baseName[idx].c < baseTyped[idx].c {
				multiByte = true
			}
			continue
		}
		return false
	}

	return multiByte
}
