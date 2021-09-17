package Solution

import "fmt"

//	暴力查找
func longestPalindrome1(s string) string {
	if len(s) == 1 {
		return s
	}
	ans := ""
	ansLength := 0

	for i := 0; i < len(s); i++ {
		for j := len(s); j > i+ansLength; j-- {
			subString := s[i:j]
			if isPalindrome(subString) && len(subString) > ansLength {
				ans = subString
				ansLength = len(subString)
			}
		}
	}

	return ans
}

func isPalindrome(s string) bool {
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}

	return true
}

// 从中心开始向2边比较
func longestPalindrome2(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		len := Max(len1, len2)

		if len > end-start {

			start = i - (len-1)/2
			end = i + len/2
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	L, R := left, right
	for L >= 0 && R < len(s) && s[L] == s[R] {
		L--
		R++
	}

	return R - L - 1
}

func Max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func Min(x, y int) int {
	if x > y {
		return y
	}

	return x
}

// Manacher
func longestPalindrome3(s string) string {
	return longestPalindromeLinear(s)
}

//	初始化Manacher需要的字符串
func initManacherStr(s string) string {
	ans := make([]rune, 0)
	ans = append(ans, '$')
	for i := 0; i < len(s); i++ {
		ans = append(ans, rune(s[i]), '$')
	}
	return string(ans)
}

//	Manacher
func longestPalindromeLinear(in string) string {
	/*
	*	初始化字符串，在字符串的没个字符左右都插入一个字符"$"
	*
	 */
	s := initManacherStr(in)
	c, max := 0, 0

	/*
	* 	P记录以每个字符为中心的最长回文串的信息
	*	P［id］记录的是以字符str［id］为中心的最长回文串
	*
	 */
	//	P记录以每个字符为中心的最长回文串的信息
	P := make([]int, len(s))
	//for i := 0; i < len(P); i++ {
	//	P[i] = 0
	//}

	for i := 1; i < len(s)-1; i++ {
		i_mirror := 2*c - i
		if max > i {
			P[i] = Min(P[i_mirror], max-i)
		} else {
			P[i] = 0
		}

		for (i+P[i]+1) < len(s) && (i-P[i]-1) >= 0 &&
			s[i+P[i]+1] == s[i-P[i]-1] {

			P[i]++
		}

		if i+P[i] > max {
			c = i
			max = i + P[i]
		}

	}
	return extractLongest(in, P)
}

func extractLongest(s string, P []int) string {
	longestCenter, longestLength := 0, 0
	for i, v := range P {
		if v > longestLength {
			longestLength = v
			longestCenter = i
		}
	}
	offset := (longestCenter - longestLength) / 2
	return s[offset : offset+longestLength]
}

//	DP
func longestPalindrome4(s string) string {
	n := len(s)
	dp := [][]bool{}
	for i := 0; i < n; i++ {
		dp = append(dp, make([]bool, n+1))
	}

	for i := 0; i < n; i++ {
		dp[i][i] = true
		if i == n-1 {
			break
		}
		dp[i][i+1] = s[i] == s[i+1]
	}

	for i := n - 3; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
		}
	}
	max := 0
	maxStr := ""
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if dp[i][j] == true && j-i+1 > max {
				max = j - i + 1
				maxStr = s[i : j+1]
			}
		}
	}

	return maxStr
}

func longestPalindrome5(s string) string {
	dp := make([][]bool, 0)
	left, right := 0, 0
	for i := 0; i <= len(s); i++ {
		dp = append(dp, make([]bool, len(s)))
	}

	for i := len(s) - 1; i >= 0; i-- {
		dp[i][i] = true
		for j := i; j < len(s); j++ {
			dp[j][i] = s[i] == s[j] && (i-j < 2 || dp[j+1][i-1])
			if dp[j][i] && j-i > right-left {
				left = i
				right = j
			}
			fmt.Println(i, j)
			Print(dp)

		}
	}
	return s[left : right+1]
}

func Print(x [][]bool) {
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}
	fmt.Println()
}
