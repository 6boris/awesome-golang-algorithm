package Solution

import (
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs int
		expect []int
	}{
		{"TestCase", 5, []int{-7, -1, 1, 3, 4}},
		{"TestCase", 4, []int{-2, -1, 1, 2}},
		{"TestCase", 3, []int{-1, 0, 1}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs)
			res := checkResult(got)
			if res != 0 {
				t.Fatalf("expected: 0, but got: %v, with inputs: %v", res, c.inputs)
			}
		})
	}
}

func checkResult(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

//	压力测试
func BenchmarkSolution(b *testing.B) {

}

//	使用案列
func ExampleSolution() {

}
