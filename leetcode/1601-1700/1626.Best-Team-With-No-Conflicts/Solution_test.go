package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name         string
		scores, ages []int
		expect       int
	}{
		{"TestCase1", []int{1, 3, 5, 10, 15}, []int{1, 2, 3, 4, 5}, 34},
		{"TestCase2", []int{4, 5, 6, 5}, []int{2, 1, 2, 1}, 16},
		{"TestCase3", []int{1, 2, 3, 5}, []int{8, 9, 10, 1}, 6},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.scores, c.ages)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.scores, c.ages)
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
