package Solution

func Solution(letters []byte, target byte) byte {
	x := uint8(0)
	for i := 0; i < len(letters); i++ {
		if letters[i] > target {
			if x == 0 || letters[i] < x {
				x = letters[i]
			}
		}
	}
	if x == 0 {
		return letters[0]
	}
	return x
}
