package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name           string
		landStartTime  []int
		landDuration   []int
		waterStartTime []int
		waterDuration  []int
		expect         int
	}{
		{"TestCase1", []int{2, 8}, []int{4, 1}, []int{6}, []int{3}, 9},
		{"TestCase2", []int{5}, []int{3}, []int{1}, []int{10}, 14},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.landStartTime, c.landDuration, c.waterStartTime, c.waterDuration)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v %v",
					c.expect, got, c.landStartTime, c.landDuration, c.waterStartTime, c.waterDuration)
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
