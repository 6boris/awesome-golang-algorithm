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
		inputs []int
		expect int
	}{
		{"TestCase1", []int{1, -2, 3, -2}, 3},
		{"TestCase2", []int{-3, -2, -3}, -2},
		{"TestCase3", []int{5, -3, 5}, 10},
		{"TestCase4", []int{1}, 1},
		{"TestCase5", []int{1, 2, 3, 4}, 10},
		{"TestCase6", []int{1, 2}, 3},
		{"TestCase7", []int{1}, 1},
		{"TestCase8", []int{1, -2, 3}, 4},
		{"TestCase9", []int{-2, 2, -2, 9}, 9},
		{"TestCase10", []int{8, -15, -29, -19}, 8},
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
