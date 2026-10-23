package Solution

func Solution(num int, sum int) string {
	if sum < 0 || sum > num*9 {
		return ""
	}
	ret := make([]byte, num)
	index := 0
	for ; sum > 9 && index < num; sum, index = sum-9, index+1 {
		ret[index] = '9'
	}
	if index >= num && sum != 0 {
		return ""
	}
	ret[index] = '0' + byte(sum)
	index++
	for ; index < num; index++ {
		ret[index] = '0'
	}
	return string(ret)
}
