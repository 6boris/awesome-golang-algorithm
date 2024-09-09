package Solution

type char809 struct {
	c  byte
	cc int
}

func toChars(s string) []char809 {
	pre := s[0]
	c := 1
	chars := make([]char809, 0)
	for i := 1; i < len(s); i++ {
		b := s[i]
		if b == pre {
			c++
			continue
		}
		chars = append(chars, char809{c: byte(pre), cc: c})
		pre = b
		c = 1
	}
	chars = append(chars, char809{c: byte(pre), cc: c})
	return chars
}
func Solution(s string, words []string) int {
	ans := 0
	target := toChars(s)
	for _, word := range words {
		tmp := toChars(word)
		if len(tmp) != len(target) {
			continue
		}
		ok := true
		for i := range len(target) {
			if target[i].c != tmp[i].c {
				ok = false
				break
			}
			if target[i].cc < 3 && tmp[i].cc != target[i].cc {
				ok = false
				break
			}
			if tmp[i].cc > target[i].cc {
				ok = false
				break
			}
		}
		if ok {
			ans++
		}
	}
	return ans
}
