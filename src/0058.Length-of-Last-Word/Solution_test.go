package Solution

import "testing"

func TestSolution1(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		got := lengthOfLastWord("Hello World")
		want := 5
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

	t.Run("Test-2", func(t *testing.T) {
		got := lengthOfLastWord("Hello World ")
		want := 5
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

}

func TestSolution2(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		got := lengthOfLastWord1("Hello World")
		want := 5
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

	t.Run("Test-2", func(t *testing.T) {
		got := lengthOfLastWord1("Hello World ")
		want := 5
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

	t.Run("Test-3", func(t *testing.T) {
		got := lengthOfLastWord1(" H ")
		want := 5
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})
}
