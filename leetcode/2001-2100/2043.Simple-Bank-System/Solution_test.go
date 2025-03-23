package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name     string
		balances []int64
		actions  []action
		expect   []bool
	}{
		{"TestCase1", []int64{10, 100, 20, 50, 30}, []action{
			{"w", 3, 0, 10},
			{"t", 5, 1, 20},
			{"d", 5, 0, 20},
			{"t", 3, 4, 15},
			{"w", 10, 0, 50},
		}, []bool{true, true, true, false, false}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.balances, c.actions)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.balances, c.actions)
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
