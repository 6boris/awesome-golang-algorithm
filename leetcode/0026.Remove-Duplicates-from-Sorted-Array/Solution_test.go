package Solution

import "testing"

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		data := []int{1, 1, 2}
		got := removeDuplicates(data)
		want := 2
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

	t.Run("Test-2", func(t *testing.T) {
		data := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
		got := removeDuplicates(data)
		want := 5
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

}
