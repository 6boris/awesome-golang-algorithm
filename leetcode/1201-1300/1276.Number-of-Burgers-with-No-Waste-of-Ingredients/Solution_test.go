package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name                       string
		tomatoSlices, cheeseSlices int
		expect                     []int
	}{
		{"TestCase1", 16, 7, []int{1, 6}},
		{"TestCase2", 17, 4, []int{}},
		{"TestCase3", 4, 17, []int{}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.tomatoSlices, c.cheeseSlices)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.tomatoSlices, c.cheeseSlices)
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
