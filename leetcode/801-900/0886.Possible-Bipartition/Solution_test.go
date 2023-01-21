package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name     string
		n        int
		dislikes [][]int
		expect   bool
	}{
		{"TestCase1", 4, [][]int{{1, 2}, {1, 3}, {2, 4}}, true},
		{"TestCase2", 3, [][]int{{1, 2}, {1, 3}, {2, 3}}, false},
		{"TestCase5", 5, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {1, 5}}, false},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.n, c.dislikes)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.n, c.dislikes)
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
