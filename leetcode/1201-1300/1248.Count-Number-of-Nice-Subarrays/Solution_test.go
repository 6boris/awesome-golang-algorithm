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
		k      int
		expect int
	}{
		{"TestCase1", []int{1, 1, 2, 1, 1}, 3, 2},
		{"TestCase2", []int{2, 4, 6}, 1, 0},
		{"TestCase3", []int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 2, 16},
		{"TestCase4", []int{91473, 45388, 24720, 35841, 29648, 77363, 86290, 58032, 53752, 87188, 34428, 85343, 19801, 73201}, 4, 6},
		{"TestCase5", []int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 1, 24},
		{"TestCase6", []int{2, 1, 1}, 1, 3},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.k)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.inputs, c.k)
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
