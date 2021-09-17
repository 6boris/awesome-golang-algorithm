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
		amount int
		coins  []int
		expect int
	}{
		{"TestCase1", 5, []int{1, 2, 5}, 4},
		{"TestCase2", 10, []int{10}, 1},
		{"TestCase3", 3, []int{2}, 0},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.amount, c.coins)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with input-amount: %v, input-coins: %v",
					c.expect, got, c.amount, c.coins)
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
