package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name          string
		inputs        []int
		first, second int
		expect        int
	}{
		{"TestCase1", []int{0, 6, 5, 2, 2, 5, 1, 9, 4}, 1, 2, 20},
		{"TestCase2", []int{3, 8, 1, 3, 2, 1, 8, 9, 0}, 3, 2, 29},
		{"TestCase3", []int{2, 1, 5, 6, 0, 9, 5, 0, 3, 8}, 4, 3, 31},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.first, c.second)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.inputs, c.first, c.second)
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
