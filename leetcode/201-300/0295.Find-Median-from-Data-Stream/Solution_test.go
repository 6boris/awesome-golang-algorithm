package Solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	var medianFinder MedianFinder
	var median float64

	// Test Case 1
	medianFinder = Constructor()
	medianFinder.AddNum(1)
	medianFinder.AddNum(2)
	median = medianFinder.FindMedian()
	if !reflect.DeepEqual(median, 1.5) {
		t.Fatalf("expected: %v, but got: %v", 1.5, median)
	}
	medianFinder.AddNum(3)
	median = medianFinder.FindMedian()
	if !reflect.DeepEqual(median, 2.0) {
		t.Fatalf("expected: %v, but got: %v", 2.0, median)
	}

	// Test Case 2
	medianFinder = Constructor()
	medianFinder.AddNum(2)
	medianFinder.AddNum(1)
	median = medianFinder.FindMedian()
	if !reflect.DeepEqual(median, 1.5) {
		t.Fatalf("expected: %v, but got: %v", 1.5, median)
	}
	medianFinder.AddNum(5)
	median = medianFinder.FindMedian()
	if !reflect.DeepEqual(median, 2.0) {
		t.Fatalf("expected: %v, but got: %v", 2.0, median)
	}
	medianFinder.AddNum(7)
	medianFinder.AddNum(2)
	median = medianFinder.FindMedian()
	if !reflect.DeepEqual(median, 2.0) {
		t.Fatalf("expected: %v, but got: %v", 2.0, median)
	}
	medianFinder.AddNum(0)
	medianFinder.AddNum(5)
	median = medianFinder.FindMedian()
	if !reflect.DeepEqual(median, 2.0) {
		t.Fatalf("expected: %v, but got: %v", 2.0, median)
	}
}

//	压力测试
func BenchmarkSolution(b *testing.B) {

}

//	使用案列
func ExampleSolution() {

}
