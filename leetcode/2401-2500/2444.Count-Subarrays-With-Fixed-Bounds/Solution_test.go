package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name       string
		inputs     []int
		minK, maxK int
		expect     int64
	}{
		{"TestCase1", []int{35054, 398719, 945315, 945315, 820417, 945315, 35054, 945315, 171832, 945315, 35054, 109750, 790964, 441974, 552913}, 35054, 945315, int64(81)},
		{"TestCase2", []int{1, 3, 5, 2, 7, 5}, 1, 5, int64(2)},
		{"TestCase3", []int{1, 1, 1, 1}, 1, 1, int64(10)},
		{"TestCase4", []int{978650, 978650, 978650, 68071, 52201, 68071, 186141, 978650, 978650, 267135, 68071, 717241, 242895, 68071, 582505, 978650, 68071, 68071}, 68071, 978650, int64(57)},
		{"TestCase5", []int{1, 4, 10, 2, 4, 1, 2, 10, 2, 4, 1, 2, 4, 1, 21, 43, 12}, 3, 13, int64(0)},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.minK, c.maxK)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
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
