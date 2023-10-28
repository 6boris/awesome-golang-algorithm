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
		nums   []int
		s      string
		d      int
		expect int
	}{
		{"TestCase1", []int{-2, 0, 2}, "RLL", 3, 8},
		{"TestCase2", []int{1, -67, 68, -26, -13, -40, -56, 62, 21}, "LLLRLLRRR", 4, 2106},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.nums, c.s, c.d)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.nums, c.s, c.d)
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
