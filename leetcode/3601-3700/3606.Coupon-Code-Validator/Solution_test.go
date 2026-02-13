package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name               string
		code, businessLine []string
		isActive           []bool
		expect             []string
	}{
		{"TestCase1", []string{"SAVE20", "", "PHARMA5", "SAVE@20"}, []string{"restaurant", "grocery", "pharmacy", "restaurant"}, []bool{true, true, true, true}, []string{"PHARMA5", "SAVE20"}},
		{"TestCase2", []string{"GROCERY15", "ELECTRONICS_50", "DISCOUNT10"}, []string{"grocery", "electronics", "invalid"}, []bool{false, true, true}, []string{"ELECTRONICS_50"}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.code, c.businessLine, c.isActive)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.code, c.businessLine, c.isActive)
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
