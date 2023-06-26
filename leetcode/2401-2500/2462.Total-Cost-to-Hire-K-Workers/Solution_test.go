package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name          string
		costs         []int
		k, candidates int
		expect        int64
	}{
		{"TestCase1", []int{17, 12, 10, 2, 7, 2, 11, 20, 8}, 3, 4, 11},
		{"TestCase2", []int{1, 2, 4, 1}, 3, 3, 4},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.costs, c.k, c.candidates)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.costs, c.k, c.candidates)
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
