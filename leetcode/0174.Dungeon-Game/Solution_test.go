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
		input1 []int
		input2 [][]int
		expect []int
	}{
		{"TestCase", []int{-2, 0, 3, -5, 2, -1}, [][]int{{0, 2}, {2, 5}, {0, 5}}, []int{1, -1, -3}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			obj := Constructor(c.input1)
			for i := 0; i < len(c.input2); i++ {
				got := obj.SumRange(c.input2[i][0], c.input2[i][1])
				if !reflect.DeepEqual(got, c.expect[i]) {
					t.Fatalf("expected: %v, but got: %v, with input2: %v input1: %v ",
						c.expect[i], got, c.input2[i], c.input1[i])
				}
			}
		})
	}
}

func TestSolution2(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		input1 []int
		input2 [][]int
		expect []int
	}{
		{"TestCase", []int{-2, 0, 3, -5, 2, -1}, [][]int{{0, 2}, {2, 5}, {0, 5}}, []int{1, -1, -3}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			obj := Constructor2(c.input1)
			for i := 0; i < len(c.input2); i++ {
				got := obj.SumRange2(c.input2[i][0], c.input2[i][1])
				if !reflect.DeepEqual(got, c.expect[i]) {
					t.Fatalf("expected: %v, but got: %v, with input2: %v input1: %v ",
						c.expect[i], got, c.input2[i], c.input1[i])
				}
			}
		})
	}
}
func TestSolution3(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		input1 []int
		input2 [][]int
		expect []int
	}{
		{"TestCase", []int{-2, 0, 3, -5, 2, -1}, [][]int{{0, 2}, {2, 5}, {0, 5}}, []int{1, -1, -3}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			obj := Constructor3(c.input1)
			for i := 0; i < len(c.input2); i++ {
				got := obj.SumRange3(c.input2[i][0], c.input2[i][1])
				if !reflect.DeepEqual(got, c.expect[i]) {
					t.Fatalf("expected: %v, but got: %v, with input2: %v input1: %v ",
						c.expect[i], got, c.input2[i], c.input1[i])
				}
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
