package Solution

func intToRoman(num int) string {
	i := [10]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	x := [10]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	c := [10]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	m := [4]string{"", "M", "MM", "MMM"}
	return m[num/1000] + c[num%1000/100] + x[num%1000%100/10] + i[num%1000%100%10]
}
