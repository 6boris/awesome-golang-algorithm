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
		classes       [][]int
		extraStudents int
		expect        float64
	}{
		{"TestCase1", [][]int{{1, 2}, {3, 5}, {2, 2}}, 2, 0.7833333333333333},
		{"TestCase2", [][]int{{2, 4}, {3, 9}, {4, 5}, {2, 10}}, 4, 0.5348484848484849},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.classes, c.extraStudents)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.classes, c.extraStudents)
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
