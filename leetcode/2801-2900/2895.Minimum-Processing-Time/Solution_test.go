package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name                 string
		processorTime, tasks []int
		expect               int
	}{
		{"TestCase1", []int{8, 10}, []int{2, 2, 3, 1, 8, 7, 4, 5}, 16},
		{"TestCase2", []int{10, 20}, []int{2, 3, 1, 2, 5, 8, 4, 3}, 23},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.processorTime, c.tasks)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.processorTime, c.tasks)
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
