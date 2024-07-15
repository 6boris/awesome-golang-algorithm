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
		{"TestCase1", []int{8, 4, 6, 2, 3}, []int{4, 2, 4, 2, 3}},
		{"TestCase2", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"TestCase3", []int{10, 1, 1, 6}, []int{9, 0, 1, 6}},
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
