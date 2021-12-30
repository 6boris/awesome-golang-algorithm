package Solution

import (
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs [][]int
		k      int
		expect [][]int
	}{
		{"TestCase1", [][]int{{1, 3}, {-2, 2}}, 1, [][]int{{-2, 2}}},
		{"TestCase2", [][]int{{3, 3}, {5, -1}, {-2, 4}}, 2, [][]int{{3, 3}, {-2, 4}}},
		{"TestCase3", [][]int{{3, 3}, {5, -1}, {-2, 4}}, 3, [][]int{{3, 3}, {-2, 4}, {5, -1}}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.k)
			sort.Slice(got, func(i, j int) bool {
				if got[i][0] == got[j][0] {
					return got[i][1] < got[j][1]
				}
				return got[i][0] < got[j][0]
			})

			sort.Slice(c.expect, func(i, j int) bool {
				if c.expect[i][0] == c.expect[j][0] {
					return c.expect[i][1] < c.expect[j][1]
				}
				return c.expect[i][0] < c.expect[j][0]
			})
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
