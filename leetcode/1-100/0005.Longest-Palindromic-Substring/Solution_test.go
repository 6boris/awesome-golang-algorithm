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

func TestSolution4(t *testing.T) {
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
			ret := longestPalindrome4(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputs)
			}
		})
	}
}

func TestSolution5(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs string
		expect string
	}{
		{"1 Test Case 1", "babab", "babab"},
		{"2 Test Case 2", "cbbd", "bb"},
		{"3 Test Case 3", "a", "a"},
		{"4 Test Case 4", "bb", "bb"},
		{"5 Test Case 5", "ccc", "ccc"},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := longestPalindrome5(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputs)
			}
		})
	}
}
