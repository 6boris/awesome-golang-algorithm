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
		target int
		expect int
	}{
		{"TestCase1", []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, 8, 20},
		{"TestCase2", []int{1, 1, 2, 2, 2, 2}, 5, 12},
		{"TestCase3", []int{2, 1, 3}, 6, 1},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.target)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.inputs, c.target)
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
