package Solution

func replaceSpace(str []byte, length int) []byte {
	count := 0
	// 遍历一遍字符串， 统计字符出现的数目, 计算替换后的字符串长度
	for i := 0; i < length; i++ {
		if str[i] == ' ' {
			count++
		}
	}
	newLength := length + count*2
	// 两个index，一个指向length-1, 另一个指向newLength-1，遍历一遍字符串，完成替换
	for l, nl := length-1, newLength-1; l >= 0 && nl >= 0; {
		if str[l] == ' ' {
			str[nl] = '0'
			nl--
			str[nl] = '2'
			nl--
			str[nl] = '%'
			nl--
			l--
		} else {
			str[nl] = str[l]
			nl--
			l--
		}
	}
	return str
}
