package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name         string
		differences  []int
		lower, upper int
		expect       int
	}{
		{"TestCase1", []int{1, -3, 4}, 1, 6, 2},
		{"TestCase2", []int{3, -4, 5, 1, -2}, -4, 5, 4},
		{"TestCase3", []int{4, -7, 2}, 3, 6, 0},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.differences, c.lower, c.upper)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.differences, c.lower, c.upper)
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
