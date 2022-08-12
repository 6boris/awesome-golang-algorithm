package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name                                 string
		buckets, minutesToDie, minutesToTest int
		expect                               int
	}{
		{"TestCase1", 1000, 15, 60, 5},
		{"TestCase2", 4, 15, 15, 2},
		{"TestCase3", 4, 15, 30, 2},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.buckets, c.minutesToDie, c.minutesToTest)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.buckets, c.minutesToDie, c.minutesToTest)
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
