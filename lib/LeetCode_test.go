package Solution

import "testing"

func TestLogin(t *testing.T) {

	t.Run("test1", func(t *testing.T) {
		got := Login()
		if got != true {
			t.Fail()
		}
	})

}
