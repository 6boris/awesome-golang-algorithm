package Solution

import "fmt"

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
	fmt.Println(strArr)
	res := make([]byte, 0)

	for _, str := range strArr {
		res = append(res, str...)
	}
	return string(res)
}
func convert2(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	lenth, cycle, n := len(s), 2*numRows-2, 0
	result := make([]byte, lenth)
	for i := 0; i < numRows; i++ {
		for j1 := i; j1 < lenth; j1 = cycle + j1 {

			n += copy(result[n:], string(s[j1]))
			j2 := j1 + cycle - 2*i
			if i != 0 && i != numRows-1 && j2 < lenth {
				n += copy(result[n:], string(s[j2]))
			}

		}
	}
	return string(result)
}
