package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name        string
		distance    []int
		start       int
		destination int
		expect      int
	}{
		{"TestCase", []int{1, 2, 3, 4}, 0, 1, 1},
		{"TestCase", []int{1, 2, 3, 4}, 0, 2, 3},
		{"TestCase", []int{1, 2, 3, 4}, 0, 3, 4},
		{"TestCase", []int{1, 2, 3, 4}, 1, 4, 1},
		{"TestCase", []int{1, 2, 3, 4}, 3, 2, 3},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.distance, c.start, c.destination)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.distance, c.start, c.destination)
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
