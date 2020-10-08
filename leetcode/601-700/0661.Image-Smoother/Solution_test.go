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
		inputs [][]int
		expect [][]int
	}{
		{"TestCase", [][]int{{2, 3, 4}, {5, 6, 7}, {8, 9, 10}, {11, 12, 13}, {14, 15, 16}}, [][]int{{4, 4, 5}, {5, 6, 6}, {8, 9, 9}, {11, 12, 12}, {13, 13, 14}}},
		{"TestCase", [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}},
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
