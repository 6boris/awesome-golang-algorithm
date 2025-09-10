package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name                   string
		n                      int
		languages, friendships [][]int
		expect                 int
	}{
		{"TestCase1", 2, [][]int{{1}, {2}, {1, 2}}, [][]int{{1, 2}, {1, 3}, {2, 3}}, 1},
		{"TestCase2", 3, [][]int{{2}, {1, 3}, {1, 2}, {3}}, [][]int{{1, 4}, {1, 2}, {3, 4}, {2, 3}}, 2},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.n, c.languages, c.friendships)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.n, c.languages, c.friendships)
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
