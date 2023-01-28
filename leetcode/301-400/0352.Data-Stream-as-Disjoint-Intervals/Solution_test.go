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
		ops    []string
		nums   []int
		expect [][][]int
	}{
		{"TestCase1", []string{"addNum", "getIntervals", "addNum", "getIntervals", "addNum", "getIntervals", "addNum", "getIntervals", "addNum", "getIntervals"}, []int{1, 0, 3, 0, 7, 0, 2, 0, 6, 0}, [][][]int{
			{{1, 1}}, {{1, 1}, {3, 3}}, {{1, 1}, {3, 3}, {7, 7}}, {{1, 3}, {7, 7}}, {{1, 3}, {6, 7}},
		}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.ops, c.nums)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.ops, c.name)
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
