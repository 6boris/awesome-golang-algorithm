package Solution

import "strings"

var (
	basic = map[int]string{
		0: "Zero", 1: "One", 2: "Two", 3: "Three", 4: "Four", 5: "Five",
		6: "Six", 7: "Seven", 8: "Eight", 9: "Nine", 10: "Ten", 11: "Eleven", 12: "Twelve",
		13: "Thirteen", 14: "Fourteen", 15: "Fifteen", 16: "Sixteen", 17: "Seventeen",
		18: "Eighteen", 19: "Nineteen", 20: "Twenty", 30: "Thirty", 40: "Forty", 50: "Fifty",
		60: "Sixty", 70: "Seventy", 80: "Eighty", 90: "Ninety",
	}
	units = []string{"", "Thousand", "Million", "Billion"}
)

func hundred2english(n, index int) []string {
	if n == 0 {
		return nil
	}

	ans := []string{units[index]}
	low := n % 100
	n /= 100
	if low != 0 {
		if v, ok := basic[low]; ok {
			ans = append(ans, v)
		} else {
			mod := low % 10
			if mod != 0 {
				ans = append(ans, basic[mod])
			}
			ans = append(ans, basic[low-mod])
		}
	}

	if n != 0 {
		ans = append(ans, "Hundred", basic[n])
	}
	return ans
}

func Solution(num int) string {
	if v, ok := basic[num]; ok {
		return v
	}

	ans := make([]string, 0)
	unitIndx := 0
	for num > 0 {
		cur := num % 1000
		ans = append(ans, hundred2english(cur, unitIndx)...)
		num /= 1000
		unitIndx++
	}
	for s, e := 0, len(ans)-1; s < e; s, e = s+1, e-1 {
		ans[s], ans[e] = ans[e], ans[s]
	}
	r := strings.Join(ans, " ")
	r = strings.TrimSpace(r)
	return r
}
