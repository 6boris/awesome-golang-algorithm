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
		inputs int
		m      *MountainArray
		expect int
	}{
		{"TestCase1", 3, &MountainArray{data: []int{1, 2, 3, 4, 5, 3, 1}}, 2},
		{"TestCase2", 3, &MountainArray{data: []int{0, 1, 2, 4, 2, 1}}, -1},
		{"TestCase3", 2, &MountainArray{data: []int{1, 2, 3, 4, 5, 3, 1}}, 1},
		{"TestCase4", 5, &MountainArray{data: []int{1, 5, 2}}, 1},
		{"TestCase5", 2, &MountainArray{data: []int{1, 5, 2}}, 2},
		{"TestCase6", 0, &MountainArray{data: []int{1, 2, 3, 5, 3}}, -1},
		{"TestCase7", 1, &MountainArray{data: []int{1, 2, 3, 5, 3}}, 0},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.m)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.inputs, c.m.data)
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
