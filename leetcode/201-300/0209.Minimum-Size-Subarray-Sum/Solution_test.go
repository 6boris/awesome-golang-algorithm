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
		target int
		nums   []int
		expect int
	}{
		{"TestCase1", 7, []int{2, 3, 1, 2, 4, 3}, 2},
		{"TestCase2", 4, []int{1, 4, 4}, 1},
		{"TestCase3", 11, []int{1, 1, 1, 1, 1, 1, 1, 1}, 0},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.target, c.nums)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.target, c.nums)
			}
		})
	}
}

//	压力测试
func BenchmarkSolution(b *testing.B) {
}

//	使用案列
func ExampleSolution() {
}
