package Solution

type tmp1106 struct {
	op byte

	f, t bool
}

func Solution(expression string) bool {
	l := len(expression)
	stack := make([]tmp1106, l)
	index := -1
	f, t := false, false
	i := 0
	for ; i < l; i++ {
		cur := expression[i]
		if cur == '!' || cur == '|' || cur == '&' {
			index++
			stack[index] = tmp1106{
				cur, f, t,
			}
			t, f = false, false
			i++
			continue
		}
		if cur == ')' {
			top := stack[index]
			index--
			if top.op == '!' {
				if t {
					t, f = false, true
					f = true
				} else {
					t, f = true, false
				}
			}
			if top.op == '|' {
				if t {
					t, f = true, false
				} else {
					t, f = false, true
				}
			}
			if top.op == '&' {
				if f {
					t, f = false, true
				} else {
					t, f = true, false
				}
			}
			t = t || top.t
			f = f || top.f
			continue
		}
		if cur == 'f' {
			f = true
		}
		if cur == 't' {
			t = true
		}
	}
	return t
}
