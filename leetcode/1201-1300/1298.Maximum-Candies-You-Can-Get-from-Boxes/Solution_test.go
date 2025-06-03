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
		status, candies      []int
		keys, containedBoxes [][]int
		initialBoxes         []int
		expect               int
	}{
		{"TestCase1", []int{1, 0, 1, 0}, []int{7, 5, 4, 100}, [][]int{{}, {}, {1}, {}}, [][]int{{1, 2}, {3}, {}, {}}, []int{0}, 16},
		{"TestCase2", []int{1, 0, 0, 0, 0, 0}, []int{1, 1, 1, 1, 1, 1}, [][]int{{1, 2, 3, 4, 5}, {}, {}, {}, {}, {}}, [][]int{{1, 2, 3, 4, 5}, {}, {}, {}, {}, {}}, []int{0}, 6},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.status, c.candies, c.keys, c.containedBoxes, c.initialBoxes)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v %v %v",
					c.expect, got, c.status, c.candies, c.keys, c.containedBoxes, c.initialBoxes)
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
