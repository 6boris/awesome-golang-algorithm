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
		inputs []*Employee
		id     int
		expect int
	}{
		{"TestCase1", []*Employee{
			{
				Id:           1,
				Importance:   5,
				Subordinates: []int{2, 3},
			},
			{Id: 2, Importance: 3},
			{Id: 3, Importance: 3},
		}, 1, 11},
		{"TestCase2", []*Employee{
			{Id: 1, Importance: 1, Subordinates: []int{5}},
			{Id: 5, Importance: -3},
		}, 5, -3},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.id)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.inputs, c.id)
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
