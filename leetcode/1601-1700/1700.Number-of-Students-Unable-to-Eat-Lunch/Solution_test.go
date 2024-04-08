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
		students, sandwiches []int
		expect               int
	}{
		{"TestCase1", []int{1, 1, 0, 0}, []int{0, 1, 0, 1}, 0},
		{"TestCase2", []int{1, 1, 1, 0, 0, 1}, []int{1, 0, 0, 0, 1, 1}, 3},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.students, c.sandwiches)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.students, c.sandwiches)
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
