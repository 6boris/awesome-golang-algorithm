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
		expect int
	}{
		{"TestCase1", []int{0, 1, 0}, 1},
		{"TestCase2", []int{0, 2, 1, 0}, 1},
		{"TestCase3", []int{0, 10, 5, 2}, 1},
		{"TestCase4", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 5, 4, 2, 1}, 8},
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
