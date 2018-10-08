package Solution

import "testing"

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		got := Solution()
		want := false
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})
	t.Run("asd", func(t *testing.T) {

	})
	t.Run("Test-2", func(t *testing.T) {
		got := Solution()
		want := true
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})
}
