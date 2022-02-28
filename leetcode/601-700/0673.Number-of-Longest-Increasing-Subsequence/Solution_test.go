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
		inputs []int
		expect int
	}{
		{"TestCase1", []int{1, 3, 5, 4, 7}, 2},
		{"TestCase2", []int{2, 2, 2, 2, 2}, 5},
		{"TestCase3", []int{1, 2, 3}, 1},
		{"TestCase4", []int{1, 3, 5, 4, 7, 7, 8, 7, 2, 10, 19, 23, 10, 101, 234, 78, 23}, 4},
		{"TestCase5", []int{1, 1, 1, 2, 2, 2, 3, 3, 3}, 27},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
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
