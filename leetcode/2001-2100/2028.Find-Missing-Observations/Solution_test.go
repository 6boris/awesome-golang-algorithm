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
		rolls   []int
		mean, n int
		expect  []int
	}{
		{"TestCase1", []int{3, 2, 4, 3}, 4, 2, []int{6, 6}},
		{"TestCase2", []int{1, 5, 6}, 3, 4, []int{3, 2, 2, 2}},
		{"TestCase3", []int{1, 2, 3, 4}, 6, 4, []int{}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.rolls, c.mean, c.n)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.rolls, c.mean, c.n)
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
