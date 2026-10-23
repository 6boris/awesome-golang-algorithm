package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name       string
		rec1, rec2 []int
		expect     bool
	}{
		{"TestCase1", []int{0, 0, 2, 2}, []int{1, 1, 3, 3}, true},
		{"TestCase2", []int{0, 0, 1, 1}, []int{1, 0, 2, 1}, false},
		{"TestCase3", []int{0, 0, 1, 1}, []int{2, 2, 3, 3}, false},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.rec1, c.rec2)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.rec1, c.rec2)
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
