package Solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []int
		expect []int
	}{
		{"TestCacse 1", []int{1, 2, 3}, []int{1, 3, 2}},
		{"TestCacse 2", []int{3, 2, 1}, []int{1, 2, 3}},
		{"TestCacse 3", []int{1, 1, 5}, []int{1, 5, 1}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			nextPermutation(c.inputs)
			if !reflect.DeepEqual(c.inputs, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, c.inputs, c.inputs)
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
