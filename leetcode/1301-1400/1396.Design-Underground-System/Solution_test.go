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
		inputs []op
		expect []float64
	}{
		{"TestCase1", []op{
			{action: "in", station: "Leyton", id: 45, t: 3},
			{action: "in", station: "Paradise", id: 32, t: 8},
			{action: "in", station: "Leyton", id: 27, t: 10},
			{action: "out", station: "Waterloo", id: 45, t: 15},
			{action: "out", station: "Waterloo", id: 27, t: 20},
			{action: "out", station: "Cambridge", id: 32, t: 22},
			{action: "avg", station: "Paradise", end: "Cambridge"},
			{action: "avg", station: "Leyton", end: "Waterloo"},
			{action: "in", station: "Leyton", id: 10, t: 24},
			{action: "avg", station: "Leyton", end: "Waterloo"},
			{action: "out", station: "Waterloo", id: 10, t: 38},
			{action: "avg", station: "Leyton", end: "Waterloo"},
		}, []float64{14.0, 11.0, 11.0, 12.0}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
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
