package Solution

import "testing"

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		got := romanToInt("III")
		want := 3
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

}
