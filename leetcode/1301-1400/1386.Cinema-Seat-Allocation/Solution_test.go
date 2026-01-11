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
		n             int
		reservedSeats [][]int
		expect        int
	}{
		{"TestCase1", 3, [][]int{{1, 2}, {1, 3}, {1, 8}, {2, 6}, {3, 1}, {3, 10}}, 4},
		{"TestCase2", 2, [][]int{{2, 1}, {1, 8}, {2, 6}}, 2},
		{"TestCase3", 4, [][]int{{4, 3}, {1, 4}, {4, 6}, {1, 7}}, 4},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.n, c.reservedSeats)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.n, c.reservedSeats)
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
