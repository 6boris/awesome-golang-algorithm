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
		code   []int
		k      int
		expect []int
	}{
		{"TestCase1", []int{5, 7, 1, 4}, 3, []int{12, 10, 16, 13}},
		{"TestCase2", []int{1, 2, 3, 4}, 0, []int{0, 0, 0, 0}},
		{"TestCase3", []int{2, 4, 9, 3}, -2, []int{12, 5, 6, 13}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.code, c.k)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.code, c.k)
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
