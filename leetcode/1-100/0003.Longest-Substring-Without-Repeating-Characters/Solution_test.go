package Solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs string
		expect int
	}{
		{"1 test 1", "abcabcbb!", 3},
		{"2 test 2", "bbbbb", 1},
		{"3 test 3", "pwwkew", 3},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := lengthOfLongestSubstring(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputs)
			}
		})
	}
}

func TestSolution2(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs string
		expect int
	}{
		{"1 test 1", "abcabcbb!", 3},
		{"2 test 2", "bbbbb", 1},
		{"3 test 3", "pwwkew", 3},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := lengthOfLongestSubstring2(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputs)
			}
		})
	}
}

const N = 1000

func BenchmarkMax1(b *testing.B) {
	//b.N = 10000
	//	测试用例
	cases := []struct {
		name   string
		inputs string
		expect int
	}{
		{"1 test 1", "abcabcbb!", 3},
		{"2 test 2", "bbbbb", 1},
		{"3 test 3", "pwwkew", 3},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring(cases[0].inputs)
		lengthOfLongestSubstring(cases[1].inputs)
		lengthOfLongestSubstring(cases[2].inputs)
	}
}

func BenchmarkMax2(b *testing.B) {
	//b.N = 10000
	//	测试用例
	cases := []struct {
		name   string
		inputs string
		expect int
	}{
		{"1 test 1", "abcabcbb!", 3},
		{"2 test 2", "bbbbb", 1},
		{"3 test 3", "pwwkew", 3},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring2(cases[0].inputs)
		lengthOfLongestSubstring2(cases[1].inputs)
		lengthOfLongestSubstring2(cases[2].inputs)
	}
}
