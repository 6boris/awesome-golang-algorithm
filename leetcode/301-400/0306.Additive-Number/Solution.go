package Solution

import "strings"

func Solution(num string) bool {
	if len(num) < 3 {
		return false
	}
	if num[0] == '0' && num[1] == '0' {
		return judge306(num, 0, 0, 1)
	}

	for end := 1; end < len(num)-1; end++ {
		if num[0] == '0' {
			if judge306(num, 0, 0, end) {
				return true
			}
			continue
		}
		for splitEnd := 0; splitEnd < end; splitEnd++ {
			if num[splitEnd+1] == '0' && splitEnd != end-1 {
				continue
			}
			if judge306(num, 0, splitEnd, end) {
				return true
			}
		}
	}
	return false
}

func judge306(num string, start, mid, end int) bool {
	for end < len(num)-1 {
		tmp := addStr(num, start, mid, end)
		nextEnd := end + len(tmp)
		if nextEnd >= len(num) {
			return false
		}
		i, j := 0, nextEnd
		for ; j > end; i, j = i+1, j-1 {
			if num[j] != tmp[i] {
				return false
			}
		}
		start, mid, end = mid+1, end, nextEnd
	}
	return true
}
func addStr(num string, start, mid, end int) string {
	buf := strings.Builder{}
	i, j := mid, end
	cf := uint8(0)
	for ; i >= start || j > mid; i, j = i-1, j-1 {
		a, b := byte('0'), byte('0')
		if i >= start {
			a = byte(num[i])
		}
		if j > mid {
			b = byte(num[j])
		}
		tmp := a - '0' + b - '0' + cf
		cf = tmp / 10
		tmp %= 10
		buf.WriteByte(tmp + '0')
	}
	if cf == 1 {
		buf.WriteByte('1')
	}
	return buf.String()
}
