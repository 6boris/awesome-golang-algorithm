package Solution

import "fmt"

func Solution(chars []byte) int {
	length := len(chars)
	if length == 0 {
		return length
	}
	nowByte, sameCount := chars[0], 1
	targetIndex := 0
	for idx := 1; idx <= length; idx++ {
		if idx < length && chars[idx] == nowByte {
			sameCount++
			continue
		}
		chars[targetIndex] = nowByte
		targetIndex++
		if idx < length {
			nowByte = chars[idx]
		}
		if sameCount == 1 {
			continue
		}
		for _, b := range fmt.Sprintf("%d", sameCount) {
			chars[targetIndex] = byte(b)
			targetIndex++
		}
		sameCount = 1
	}
	return targetIndex
}
