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
		expect [][]int
	}{
		{"TestCase", "abbxxxxzzy", [][]int{{3, 6}}},
		{"TestCase", "abcdddeeeeaabbbcd", [][]int{{3, 5}, {6, 9}, {12, 14}}},
		{"TestCase", "ab", [][]int{}},
		{"TestCase", "xxxaaaa", [][]int{{0, 2}, {3, 6}}},
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

//	压力测试
func BenchmarkSolution(b *testing.B) {
}

//	使用案列
func ExampleSolution() {
}
