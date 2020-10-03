package Solution

import "testing"

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		nums := []int{3, 2, 2, 3}
		val := 3
		got := removeElement(nums, val)
		want := 2
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

	t.Run("Test-2", func(t *testing.T) {
		nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
		val := 2
		got := removeElement(nums, val)
		want := 5
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

}
