package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name           string
		servers, tasks []int
		expect         []int
	}{
		{"TestCase1", []int{3, 3, 2}, []int{1, 2, 3, 2, 1, 2}, []int{2, 2, 0, 2, 1, 2}},
		{"TestCase2", []int{5, 1, 4, 3, 2}, []int{2, 1, 2, 4, 5, 2, 1}, []int{1, 4, 1, 4, 1, 3, 2}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.servers, c.tasks)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.servers, c.tasks)
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
