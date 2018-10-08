package Solution

import (
	"fmt"
	"testing"
)

func IsEqualForSlice(a, b []int) bool {
	if len(a) != len(b) {
		fmt.Println("a")
		return false
	}
	//	为了和reflect.DeepEqual的结果保持一致：
	// 	[]int{} != []int(nil)
	if (a == nil) != (b == nil) {
		fmt.Println("a")

		return false
	}

	// 	此处的处的bounds check能够明确保证v != b[i]中的b[i]
	//	会出现越界错误，从而避免了b[i]中的越界检查从而提高效率
	//	https://go101.org/article/bounds-check-elimination.html
	b = b[:len(a)]

	//	循环检测每个元素是否相等
	for i, v := range a {
		if v != a[i] {
			return false
		}
	}
	return true
}

func TestTwoSum(t *testing.T) {

	t.Run("Test-1", func(t *testing.T) {
		data := []int{3, 2, 4}
		target := 6
		want := []int{1, 2}

		got := twoSum(data, target)
		if !IsEqualForSlice(got, want) {
			t.Error("GOT:", got, " WANT:", want)
		}
	})

	t.Run("Test-2", func(t *testing.T) {
		data := []int{2, 7, 11, 15}
		target := 9
		want := []int{0, 1}

		got := twoSum(data, target)
		if !IsEqualForSlice(got, want) {
			t.Error("GOT:", got, " WANT:", want)
		}
	})

	t.Run("Test-3", func(t *testing.T) {
		data := []int{7, 6, 5, 3, 2, 1, 4, 9, 10}
		target := 17
		want := []int{0, 8}

		got := twoSum(data, target)
		if !IsEqualForSlice(got, want) {
			t.Error("GOT:", got, " WANT:", want)
		}
	})

}

func TestTwoSum1(t *testing.T) {
	data := []int{7, 6, 5, 3, 2, 1, 4, 9, 10}
	target := 17
	fmt.Println(twoSum1(data, target))
}
