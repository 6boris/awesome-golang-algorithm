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
		maxSize int
		opts    []op
		expect  []int
	}{
		{"TestCase1", 3, []op{
			{"push", 1, 0}, {"push", 2, 0},
			{"pop", 0, 0},

			{"push", 2, 0}, {"push", 3, 0}, {"push", 4, 0},
			{"i", 5, 100}, {"i", 2, 100},
			{"pop", 0, 0},
			{"pop", 0, 0},
			{"pop", 0, 0},
			{"pop", 0, 0},
		}, []int{2, 103, 202, 201, -1}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.maxSize, c.opts)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.maxSize, c.opts)
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
