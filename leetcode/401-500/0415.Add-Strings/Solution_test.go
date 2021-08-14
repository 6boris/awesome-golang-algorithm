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
		num1, num2 string
		expect     string
	}{
		{"TestCase1", "11", "123", "134"},
		{"TestCase2", "456", "77", "533"},
		{"TestCase3", "0", "0", "0"},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.num1, c.num2)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with input-num1: %v, input-num2: %v",
					c.expect, got, c.num1, c.num2)
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
