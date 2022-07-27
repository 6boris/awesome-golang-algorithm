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
		input  []int
		expect int
	}{
		{"TestCase1", []int{1, 2, 3, 4}, 6},
		{"TestCase2", []int{3, 4, 5, 2}, 12},
		{"TestCase3", []int{3, 7, 9}, 48},
		{"TestCase4", []int{1, 2, 3, 4, 5}, 12},
		{"TestCase5", []int{1, 5, 4, 5}, 16},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.input)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.input)
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
