package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name           string
		p1, p2, p3, p4 []int
		expect         bool
	}{
		{"TestCase1", []int{0, 0}, []int{1, 1}, []int{1, 0}, []int{0, 1}, true},
		{"TestCase2", []int{0, 0}, []int{1, 1}, []int{1, 0}, []int{0, 12}, false},
		{"TestCase3", []int{1, 0}, []int{-1, 0}, []int{0, 1}, []int{0, -1}, true},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.p1, c.p2, c.p3, c.p4)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v %v",
					c.expect, got, c.p1, c.p2, c.p3, c.p4)
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
