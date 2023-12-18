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
		inputs int
		cost   []int
		expect int
	}{
		{"TestCase1", 7, []int{1, 5, 2, 2, 3, 3, 1}, 6},
		{"TestCase2", 3, []int{5, 3, 3}, 0},
		{"TestCase3", 15, []int{1, 5, 2, 2, 3, 3, 1, 1, 2, 3, 4, 5, 6, 7, 8}, 8},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.cost)
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
