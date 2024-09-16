package Solution

import "strings"

func maskEmail(email string) string {
	buf := strings.Builder{}
	cur := email[0]
	if cur >= 'A' && cur <= 'Z' {
		cur += uint8(32)
	}
	buf.WriteByte(cur)
	buf.WriteString("*****")
	i := 1
	for ; i < len(email) && email[i] != '@'; i++ {
	}
	cur = email[i-1]
	if cur >= 'A' && cur <= 'Z' {
		cur += uint8(32)
	}
	buf.WriteByte(cur)
	buf.WriteString(strings.ToLower(email[i:]))
	return buf.String()
}

func maskPhone(phone string) string {
	buf := make([]byte, 0)
	i := len(phone) - 1

	for ; i >= 0; i-- {
		if phone[i] >= '0' && phone[i] <= '9' {
			buf = append(buf, phone[i])
		}
	}
	ans := strings.Builder{}
	diff := len(buf) - 10
	if r := diff; r > 0 {
		ans.WriteByte('+')
		for ; r > 0; r-- {
			ans.WriteByte('*')
		}
		ans.WriteByte('-')
	}
	// 1, 2, 3, 4-5, 6, 7-8, 9, 10-11, 12,13
	index := len(buf) - diff - 1
	count := 3
	for ; index > 3; index-- {
		ans.WriteByte('*')
		count--
		if count == 0 {
			ans.WriteByte('-')
			count = 3
		}
	}
	for ; index >= 0; index-- {
		ans.WriteByte(buf[index])
	}
	return ans.String()
}

func Solution(s string) string {
	if strings.Contains(s, "@") {
		return maskEmail(s)
	}
	return maskPhone(s)
}
