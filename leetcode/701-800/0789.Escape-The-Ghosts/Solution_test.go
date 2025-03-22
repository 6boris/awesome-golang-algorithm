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
		ghosts [][]int
		target []int
		expect bool
	}{
		{"TestCase1", [][]int{{1, 0}, {0, 3}}, []int{0, 1}, true},
		{"TestCase2", [][]int{{1, 0}}, []int{2, 0}, false},
		{"TestCase3", [][]int{{2, 0}}, []int{1, 0}, false},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.ghosts, c.target)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.ghosts, c.target)
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
