package Solution

import (
	"math"
	"strconv"
)

func myAtoi(str string) int {
	var i, num int
	sign := 1
	if str == "" {
		return 0
	}
	//	去除前面空格
	for i < len(str) && str[i] == 32 {
		i++
	}

	if i >= len(str) {
		return 0
	}

	// 判断符号
	//	+
	if str[i] == 43 {
		i++
		//	-
	} else if str[i] == 45 {
		sign = -1
		i++
	}
	//	依次将每个字符转换为数字
	for ; i < len(str); i++ {
		//	后面为字符时，直接返回前面的数字
		if str[i] != 0 && (str[i] < 48 || str[i] > 57) {
			return num * sign
		}
		//	获得每次计算的结果
		n, _ := strconv.Atoi(string(str[i]))
		num = num*10 + n
		//	检查越界【越界就返回极限值】
		if num*sign < math.MinInt32 {
			return math.MinInt32
		} else if num*sign > math.MaxInt32 {
			return math.MaxInt32
		}
	}
	return num * sign
}
