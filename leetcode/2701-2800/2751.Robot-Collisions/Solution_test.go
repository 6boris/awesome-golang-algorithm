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
		p, h   []int
		d      string
		expect []int
	}{
		{"TestCase1", []int{5, 4, 3, 2, 1}, []int{2, 17, 9, 15, 10}, "RRRRR", []int{2, 17, 9, 15, 10}},
		{"TestCase2", []int{3, 5, 2, 6}, []int{10, 10, 15, 12}, "RLRL", []int{14}},
		{"TestCase3", []int{1, 2, 5, 6}, []int{10, 10, 11, 11}, "RLRL", []int{}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.p, c.h, c.d)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.p, c.h, c.d)
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
