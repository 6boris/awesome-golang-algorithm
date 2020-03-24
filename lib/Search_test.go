package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []int
		target int
		expect int
	}{
		{"TestCase", []int{1, 2, 2, 2, 3}, 2, 2},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := BinarySearch(c.inputs, c.target)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		})
	}
}

func TestBinarySearchLeftBound(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []int
		target int
		expect int
	}{
		{"TestCase", []int{1, 2, 2, 2, 3}, 2, 1},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := BinarySearchLeftBound(c.inputs, c.target)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		})
	}
}

func TestBinarySearchRightBound(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs []int
		target int
		expect int
	}{
		{"TestCase", []int{1, 2, 2, 2, 3}, 2, 3},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := BinarySearchRightBound(c.inputs, c.target)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		})
	}
}
