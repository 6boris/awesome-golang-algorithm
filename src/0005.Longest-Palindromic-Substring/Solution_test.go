package Solution

import (
	"reflect"
	"testing"
)

// 中间向两边找
func TestSolution1(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs string
		expect string
	}{
		{"1 test 1", "babad", "bab"},
		{"2 test 2", "cbbd", "bb"},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := longestPalindrome1(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputs)
			}
		})
	}
}

//
////	DP
//func TestSolution2(t *testing.T) {
//	//	测试用例
//	cases := []struct {
//		name   string
//		inputs string
//		expect string
//	}{
//		{"1 test 1", "babad", "bab"},
//		{"2 test 2", "cbbd", "bb"},
//	}
//
//	//	开始测试
//	for _, c := range cases {
//		t.Run(c.name, func(t *testing.T) {
//			ret := longestPalindrome2(c.inputs)
//			if !reflect.DeepEqual(ret, c.expect) {
//				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
//					c.expect, ret, c.inputs)
//			}
//		})
//	}
//}

//	Manacher
func TestSolution3(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs string
		expect string
	}{
		{"1 test 1", "babad", "bab"},
		{"2 test 2", "cbbd", "bb"},
		{"2 test 2", "a", "a"},
		{"2 test 2", "bb", "bb"},
		{"2 test 2", "ccc", "ccc"},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := longestPalindrome3(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputs)
			}
		})
	}
}
