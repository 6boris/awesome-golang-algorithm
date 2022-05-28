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
		strs        []string
		zeros, ones int
		expect      int
	}{
		{"TestCase1", []string{"10", "0001", "111001", "1", "0"}, 5, 3, 4},
		{"TestCase2", []string{"10", "0", "1"}, 1, 1, 2},
		{"TestCase3", []string{"001", "110", "0000", "0000"}, 9, 2, 3},
		{"TestCase4", []string{"10", "0001", "111001", "1", "0"}, 4, 3, 3},
		{"TestCase5", []string{"10", "1", "1", "0"}, 2, 2, 3},
		{"TestCase6", []string{"00", "000"}, 1, 10, 0},
		{"TestCase7", []string{"11", "111"}, 10, 1, 0},
		{"TestCase8", []string{"0", "0", "1", "1"}, 2, 2, 4},
		{"TestCase9", []string{"00101011"}, 36, 39, 1},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.strs, c.zeros, c.ones)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v, zeros: %v, ones: %v",
					c.expect, got, c.strs, c.zeros, c.ones)
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
