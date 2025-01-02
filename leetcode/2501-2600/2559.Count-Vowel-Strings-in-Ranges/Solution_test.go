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
		words   []string
		queries [][]int
		expect  []int
	}{
		{"TestCase1", []string{"aba", "bcb", "ece", "aa", "ee"}, [][]int{{0, 2}, {1, 4}, {1, 1}}, []int{2, 3, 0}},
		{"TestCase2", []string{"a", "e", "i"}, [][]int{{0, 2}, {0, 1}, {2, 2}}, []int{3, 2, 1}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.words, c.queries)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.words, c.queries)
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
