package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name      string
		q, qr, qg int
		expect    float64
	}{
		{"TestCase1", 1, 1, 1, 0},
		{"TestCase2", 2, 1, 1, 0.5},
		{"TestCase3", 4, 2, 2, 0.25},
		{"TestCase4", 100000009, 33, 17, 1.0},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.q, c.qr, c.qg)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.q, c.qr, c.qg)
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
