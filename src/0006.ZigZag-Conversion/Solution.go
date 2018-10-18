package Solution

func convert(s string, numRows int) string {
	if numRows == 1 || numRows <= 0 || numRows > len(s) {
		return s
	}
	strArr := make([][]byte, numRows)
	j := 0
	increasing := true

	for i := 0; i < len(s); i++ {
		strArr[j] = append(strArr[j], s[i])

		if increasing {
			j++
			if j == numRows-1 {
				increasing = false
			}
		} else {
			j--
			if j == 0 {
				increasing = true
			}
		}
	}
	res := make([]byte, 0)

	for _, str := range strArr {
		res = append(res, str...)
	}
	return string(res)
}
