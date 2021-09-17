package Solution

import "testing"

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		got := searchInsert([]int{1, 3, 5, 6}, 5)
		want := 2
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

	t.Run("Test-2", func(t *testing.T) {
		got := searchInsert([]int{1, 3, 5, 6}, 2)
		want := 1
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

	t.Run("Test-3", func(t *testing.T) {
		got := searchInsert([]int{1, 3, 5, 6}, 7)
		want := 4
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

	t.Run("Test-4", func(t *testing.T) {
		got := searchInsert([]int{1, 3, 5, 6}, 0)
		want := 0
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})
}
