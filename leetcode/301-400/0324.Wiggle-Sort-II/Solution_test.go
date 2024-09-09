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
		expect []int
	}{
		{"TestCase1", []int{1, 5, 1, 1, 6, 4}, []int{1, 6, 1, 5, 1, 4}},
		{"TestCase2", []int{1, 3, 2, 2, 3, 1}, []int{2, 3, 1, 3, 1, 2}},
		{"TestCase3", []int{1, 2, 3, 4, 5}, []int{3, 5, 2, 4, 1}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			Solution(c.inputs)
			if !reflect.DeepEqual(c.inputs, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, c.inputs, c.inputs)
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
