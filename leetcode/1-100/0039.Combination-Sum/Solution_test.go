package Solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs [][]int
		expect [][]int
	}{
		{
			"TestCacse 1",
			[][]int{{2, 3, 6, 7}, {7}},
			[][]int{{7}, {2, 2, 3}},
		},
		{
			"TestCacse 2",
			[][]int{{2, 3, 5}, {8}},
			[][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := combinationSum(c.inputs[0], c.inputs[1][0])

			if !IsEuqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		})
	}
}

//	检测结果
func IsEuqual(x, y [][]int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := 0; i < len(y); i++ {
		flag := false
		for j := 0; j < len(y); j++ {
			if reflect.DeepEqual(y[i], x[j]) {
				flag = true
			}
		}
		if flag == false {
			return false
		}
	}
	return true
}

//	压力测试
func BenchmarkSolution(b *testing.B) {
}

//	使用案列
func ExampleSolution() {
}
