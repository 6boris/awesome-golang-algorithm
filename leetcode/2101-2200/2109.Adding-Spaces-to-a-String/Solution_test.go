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
		inputs string
		spaces []int
		expect string
	}{
		{"TestCase1", "LeetcodeHelpsMeLearn", []int{8, 13, 15}, "Leetcode Helps Me Learn"},
		{"TestCase2", "icodeinpython", []int{1, 5, 7, 9}, "i code in py thon"},
		{"TestCase3", "spacing", []int{0, 1, 2, 3, 4, 5, 6}, " s p a c i n g"},
		{"TestCase4", "a", []int{0}, " a"},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.spaces)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.inputs, c.spaces)
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
