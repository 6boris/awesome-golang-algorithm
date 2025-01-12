package Solution

func Solution(s string, locked string) bool {
	l := len(s)
	if l&1 == 1 {
		return false
	}
	stack1, stack2 := make([]int, l), make([]int, l)
	i1, i2 := -1, -1
	for i, b := range s {
		if locked[i] == '0' {
			i1++
			stack1[i1] = i
			continue
		}
		if b == '(' {
			i2++
			stack2[i2] = i
			continue
		}
		if i2 != -1 {
			i2--
			continue
		}
		if i1 != -1 {
			i1--
			continue
		}
		return false
	}
	for i2 >= 0 && i1 >= 0 {
		if stack2[i2] > stack1[i1] {
			break
		}
		i2, i1 = i2-1, i1-1
	}
	return i2 == -1
}
