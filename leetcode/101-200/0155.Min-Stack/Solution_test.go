package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name    string
		actions []string
		inputs  []int
		expect  []int
	}{
		{"TestCase1", []string{push, push, push, getMin, pop, top, getMin}, []int{-2, 0, -3, 0, 0, 0, 0}, []int{-3, 0, -2}},
		{"TestCase2", []string{push, push, getMin, top, pop, push, getMin, top}, []int{1, 2, 0, 0, 0, -2, 0, 0}, []int{1, 2, -2, -2}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.actions, c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with input-actions: %v, input-inputs: %v",
					c.expect, got, c.actions, c.inputs)
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
