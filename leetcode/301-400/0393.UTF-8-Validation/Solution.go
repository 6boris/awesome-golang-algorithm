package Solution

func Solution(data []int) bool {
	l := len(data)
	shiftCmp := [4]int{0, 0b00000110, 0b00001110, 0b00011110}
	count := 0
	for idx := l - 1; idx >= 0; idx-- {
		if data[idx]&0b10000000 == 0 {
			if count > 0 {
				return false
			}
			continue
		}

		shift := data[idx] >> 6
		if shift == 0b00000010 {
			count++
			if count > 3 {
				return false
			}
			continue
		}

		if count == 0 {
			return false
		}

		shift = 6 - count
		head := data[idx] >> shift
		if head == shiftCmp[count] {
			count = 0
			continue
		}
		return false
	}
	return count == 0
}
