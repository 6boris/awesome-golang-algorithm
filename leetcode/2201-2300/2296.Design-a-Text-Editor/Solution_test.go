package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []opt
		expect []intOrStr
	}{
		{"TestCase1", []opt{
			{name: "a", intOrStr: intOrStr{b: "leetcode"}},
			{name: "d", intOrStr: intOrStr{a: 4}},
			{name: "a", intOrStr: intOrStr{b: "practice"}},
			{name: "r", intOrStr: intOrStr{a: 3}},
			{name: "l", intOrStr: intOrStr{a: 8}},
			{name: "d", intOrStr: intOrStr{a: 10}},
			{name: "l", intOrStr: intOrStr{a: 2}},
			{name: "r", intOrStr: intOrStr{a: 6}},
		}, []intOrStr{
			{a: 4},
			{b: "etpractice"},
			{b: "leet"},
			{a: 4},
			{b: ""},
			{b: "practi"},
		}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		})
	}
}

// 压力测试
func BenchmarkSolution(b *testing.B) {
}

// 使用案列
func ExampleSolution() {
}
