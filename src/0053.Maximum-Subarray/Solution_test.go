package Solution

import "testing"

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		got := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
		want := 6
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

}
