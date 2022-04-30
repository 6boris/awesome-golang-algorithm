package Solution

import "fmt"

/*
func countOf1(n int) int {
	cnt := 0
	for n > 0 {
		cnt++
		n = n & (n-1)
	}
	return cnt
}

func gen() {
	r := make(map[int][]string)
	for n := 1; n < 60 ; n++ {
		c := countOf1(n)
		if _, ok := r[c]; !ok {
			r[c] = make([]string, 0)
		}
		r[c] = append(r[c], fmt.Sprintf("%d", n))
	}
	for k, v := range r {
		fmt.Printf(`%d: '"%s"'`, k, strings.Join(v, `","`))
	}
}
*/
var (
	hour = [4][]string{
		{"0"},
		{"1", "2", "4", "8"},
		{"3", "5", "6", "9", "10"},
		{"7", "11"},
	}

	minute = [6][]string{
		{"00"},
		{"01", "02", "04", "08", "16", "32"},
		{"03", "05", "06", "09", "10", "12", "17", "18", "20", "24", "33", "34", "36", "40", "48"},
		{"07", "11", "13", "14", "19", "21", "22", "25", "26", "28", "35", "37", "38", "41", "42", "44", "49", "50", "52", "56"},
		{"15", "23", "27", "29", "30", "39", "43", "45", "46", "51", "53", "54", "57", "58"},
		{"31", "47", "55", "59"},
	}
)

func Solution(turnedOn int) []string {
	ans := make([]string, 0)
	for top := 0; top <= turnedOn && top < 4; top++ {
		binaryWatch(top, turnedOn-top, &ans)
	}
	return ans
}

func binaryWatch(top, down int, ans *[]string) {
	if down >= 6 {
		return
	}
	for _, h := range hour[top] {
		for _, m := range minute[down] {
			*ans = append(*ans, fmt.Sprintf("%s:%s", h, m))
		}
	}
}
