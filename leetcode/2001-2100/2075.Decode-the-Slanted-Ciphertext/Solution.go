package Solution

import "bytes"

func Solution(encodedText string, rows int) string {
	if rows == 1 || encodedText == "" {
		return encodedText
	}
	size := len(encodedText)
	cols := size / rows
	var buf bytes.Buffer
	for i := 0; i <= cols-rows+1; i++ {
		for j := i; j < size; j += cols + 1 {
			buf.WriteByte(encodedText[j])
		}
	}

	ret := buf.String()
	end := len(ret) - 1
	for ; ret[end] == ' '; end-- {

	}
	return ret[:end+1]
}
